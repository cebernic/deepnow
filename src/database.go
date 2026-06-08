// database.go - 数据库引擎与状态管理
// by szyok
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sync"

	_ "modernc.org/sqlite"
)

var sysDB, keyDB, ragDB *sql.DB
var authKeys sync.Map

type ModelConfig struct {
	Provider  string
	ApiType   string
	ApiKey    string
	ModelName string
	IpPort    string
}

func initDatabases() {
	sysDB, _ = sql.Open("sqlite", "system.db")
	sysDB.Exec(`CREATE TABLE IF NOT EXISTS admin_users (username TEXT PRIMARY KEY, password TEXT)`)
	sysDB.Exec(`CREATE TABLE IF NOT EXISTS model_configs (provider TEXT, api_type TEXT, api_key TEXT, model_name TEXT, ip_port TEXT, PRIMARY KEY(provider, api_type))`)
	sysDB.Exec(`CREATE TABLE IF NOT EXISTS system_rules (key TEXT PRIMARY KEY, value TEXT)`)

	var c int
	sysDB.QueryRow(`SELECT count(*) FROM admin_users`).Scan(&c)
	if c == 0 {
		sysDB.Exec(`INSERT INTO admin_users VALUES ('admin', '12345')`)
	}

	keyDB, _ = sql.Open("sqlite", "key.db")
	keyDB.Exec(`CREATE TABLE IF NOT EXISTS access_keys (api_key TEXT PRIMARY KEY, used INTEGER, created_at DATETIME)`)
	rows, _ := keyDB.Query(`SELECT api_key, used FROM access_keys`)
	for rows.Next() {
		var k string
		var u int
		rows.Scan(&k, &u)
		authKeys.Store(k, u)
	}
	rows.Close()

	os.MkdirAll("knowledge_base", 0755)
	ragDB, _ = sql.Open("sqlite", "rag.db")
	ragDB.Exec(`CREATE TABLE IF NOT EXISTS knowledge_index (id INTEGER PRIMARY KEY, chunk TEXT, source_file TEXT, embedding_json TEXT)`)
	ragDB.Exec(`ALTER TABLE knowledge_index ADD COLUMN embedding_json TEXT`)

	autoActivateModels()
}

func checkModelValid(apiType, provider string) bool {
	if provider == "" || provider == "none" {
		return false
	}
	var k, m string
	err := sysDB.QueryRow(`SELECT api_key, model_name FROM model_configs WHERE api_type=? AND provider=?`, apiType, provider).Scan(&k, &m)
	return err == nil && k != "" && m != ""
}

func autoActivateModels() {
	fmt.Println("--- [算力挂载状态] ---")

	var activePri string
	errPri := sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='active_lm_pri'`).Scan(&activePri)

	if errPri != nil || activePri == "" {
		sysDB.QueryRow(`SELECT provider FROM model_configs WHERE api_type='LM' AND api_key!='' AND model_name!='' LIMIT 1`).Scan(&activePri)
		if activePri != "" {
			sysDB.Exec(`INSERT INTO system_rules (key, value) VALUES ('active_lm_pri', ?)`, activePri)
		}
	} else if activePri != "none" && !checkModelValid("LM", activePri) {
		activePri = ""
		sysDB.QueryRow(`SELECT provider FROM model_configs WHERE api_type='LM' AND api_key!='' AND model_name!='' LIMIT 1`).Scan(&activePri)
		if activePri != "" {
			sysDB.Exec(`UPDATE system_rules SET value=? WHERE key='active_lm_pri'`, activePri)
		} else {
			sysDB.Exec(`UPDATE system_rules SET value='none' WHERE key='active_lm_pri'`)
		}
	}

	if activePri != "" && activePri != "none" {
		var mName string
		sysDB.QueryRow(`SELECT model_name FROM model_configs WHERE api_type='LM' AND provider=?`, activePri).Scan(&mName)
		fmt.Printf("✅ 大模型 [%s] (%s) (Primary) 现在启用\n", activePri, mName)
	} else {
		fmt.Println("❌ 当前无可用的大模型 (Primary LM)！引擎无法工作。")
	}

	var activeSec string
	errSec := sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='active_lm_sec'`).Scan(&activeSec)
	if errSec == nil && activeSec != "none" && !checkModelValid("LM", activeSec) {
		sysDB.Exec(`UPDATE system_rules SET value='none' WHERE key='active_lm_sec'`)
		activeSec = "none"
	}
	if activeSec != "" && activeSec != "none" {
		var mName string
		sysDB.QueryRow(`SELECT model_name FROM model_configs WHERE api_type='LM' AND provider=?`, activeSec).Scan(&mName)
		fmt.Printf("✅ 大模型 [%s] (%s) (Slave) 现在启用\n", activeSec, mName)
	}

	var activeRAG string
	errRag := sysDB.QueryRow(`SELECT value FROM system_rules WHERE key='active_embedding'`).Scan(&activeRAG)

	if errRag != nil || activeRAG == "" {
		sysDB.QueryRow(`SELECT provider FROM model_configs WHERE api_type='Embedding' AND api_key!='' AND model_name!='' LIMIT 1`).Scan(&activeRAG)
		if activeRAG != "" {
			sysDB.Exec(`INSERT INTO system_rules (key, value) VALUES ('active_embedding', ?)`, activeRAG)
		}
	} else if activeRAG != "none" && !checkModelValid("Embedding", activeRAG) {
		activeRAG = ""
		sysDB.QueryRow(`SELECT provider FROM model_configs WHERE api_type='Embedding' AND api_key!='' AND model_name!='' LIMIT 1`).Scan(&activeRAG)
		if activeRAG != "" {
			sysDB.Exec(`UPDATE system_rules SET value=? WHERE key='active_embedding'`, activeRAG)
		} else {
			sysDB.Exec(`UPDATE system_rules SET value='none' WHERE key='active_embedding'`)
		}
	}

	if activeRAG != "" && activeRAG != "none" {
		var mName string
		sysDB.QueryRow(`SELECT model_name FROM model_configs WHERE api_type='Embedding' AND provider=?`, activeRAG).Scan(&mName)
		fmt.Printf("✅ Embedding 模型 [%s] (%s) 现在启用\n", activeRAG, mName)
	} else {
		sysDB.Exec(`UPDATE system_rules SET value='0' WHERE key='rag_slider'`)
		fmt.Println("⚠️ Embedding 模型现在停用 (纯大模型模式，底层规则已强制归零)")
	}
	fmt.Println("--------------------------")
}

func getActiveModel(apiType string, role string) *ModelConfig {
	var provider string
	keyStr := "active_lm_pri"
	if apiType == "LM" && role == "slave" {
		keyStr = "active_lm_sec"
	}
	if apiType == "Embedding" {
		keyStr = "active_embedding"
	}

	err := sysDB.QueryRow(`SELECT value FROM system_rules WHERE key=?`, keyStr).Scan(&provider)
	if err != nil || provider == "" || provider == "none" {
		return nil
	}

	var conf ModelConfig
	err = sysDB.QueryRow(`SELECT provider, api_type, api_key, model_name, ip_port FROM model_configs WHERE provider=? AND api_type=?`, provider, apiType).Scan(
		&conf.Provider, &conf.ApiType, &conf.ApiKey, &conf.ModelName, &conf.IpPort)
	if err != nil || conf.ApiKey == "" || conf.ModelName == "" {
		return nil
	}
	return &conf
}

func cosineSimilarity(a, b []float64) float64 {
	var dotProduct, normA, normB float64
	for i := 0; i < len(a) && i < len(b); i++ {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

func searchKnowledge(query string) (string, float64) {
	activeRAG := getActiveModel("Embedding", "")
	if activeRAG == nil {
		return "", 0
	}

	queryVecJSON, err := generateEmbedding(query, activeRAG)
	if err != nil {
		return "", 0
	}

	var queryVec []float64
	json.Unmarshal([]byte(queryVecJSON), &queryVec)

	rows, err := ragDB.Query(`SELECT chunk, embedding_json FROM knowledge_index WHERE embedding_json IS NOT NULL`)
	if err != nil {
		return "", 0
	}
	defer rows.Close()

	var bestChunk string
	var maxSim float64 = -1.0

	for rows.Next() {
		var chunk, embJSON string
		rows.Scan(&chunk, &embJSON)

		var chunkVec []float64
		if json.Unmarshal([]byte(embJSON), &chunkVec) == nil {
			sim := cosineSimilarity(queryVec, chunkVec)
			if sim > maxSim {
				maxSim = sim
				bestChunk = chunk
			}
		}
	}
	return bestChunk, maxSim
}
