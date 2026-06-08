// main.go - 主创建入口
// by szyok
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

type ServerConfig struct {
	EndpointIP    string `json:"endpoint_ip"`
	EndpointPort  string `json:"endpoint_port"`
	WebUIIP       string `json:"webui_ip"`
	WebUIPort     string `json:"webui_port"`
	DomainName    string `json:"domain_name"`
	MaxTotalConns int    `json:"max_total_conns"`
	MaxConnsPerIP int    `json:"max_conns_per_ip"`
	CertFile      string `json:"cert_file"`
	KeyFile       string `json:"key_file"`
}

var config ServerConfig
var configFile = "aibase_server_config.json"

func loadConfig() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		config = ServerConfig{
			EndpointIP: "0.0.0.0", EndpointPort: "8443",
			WebUIIP: "0.0.0.0", WebUIPort: "8080",
			DomainName: "", MaxTotalConns: 1000, MaxConnsPerIP: 50,
			CertFile: "server.crt", KeyFile: "server.key",
		}
		saveConfig()
	} else {
		file, _ := os.ReadFile(configFile)
		json.Unmarshal(file, &config)
		needSave := false
		if config.CertFile == "" {
			config.CertFile = "server.crt"
			needSave = true
		}
		if config.KeyFile == "" {
			config.KeyFile = "server.key"
			needSave = true
		}
		if needSave {
			saveConfig()
		}
	}
}

func saveConfig() {
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configFile, data, 0644)
}

func ensureCert() {
	if _, err := os.Stat(config.CertFile); os.IsNotExist(err) {
		fmt.Println("🛡️ 未检测到外部 HTTPS 证书，正在生成兜底的高强度自签名证书...")
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)
		host := config.DomainName
		if host == "" {
			host = "localhost"
		}
		template := x509.Certificate{
			SerialNumber: big.NewInt(1),
			Subject:      pkix.Name{Organization: []string{"AI Base Server"}},
			DNSNames:     []string{host},
			NotBefore:    time.Now(), NotAfter: time.Now().Add(3650 * 24 * time.Hour),
			KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			BasicConstraintsValid: true,
		}
		derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
		certOut, _ := os.Create(config.CertFile)
		pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
		certOut.Close()
		keyOut, _ := os.Create(config.KeyFile)
		pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
		keyOut.Close()
		fmt.Println("✅ 默认自签名证书生成完毕！")
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("🚀 AI Base Server (Deepnow) (算力聚合基座) 启动中...")
	fmt.Println("==================================================")
	loadConfig()
	ensureCert()
	initDatabases()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/v1/chat/completions", handleGatewayCompletions)
		addr := fmt.Sprintf("%s:%s", config.EndpointIP, config.EndpointPort)
		fmt.Printf("🔒 Endpoint (HTTPS) -> https://%s/v1/chat/completions\n", addr)
		srv := &http.Server{Addr: addr, Handler: mux, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}}
		log.Fatal(srv.ListenAndServeTLS(config.CertFile, config.KeyFile))
	}()

	muxWeb := http.NewServeMux()
	muxWeb.Handle("/", http.FileServer(http.Dir("./static")))
	muxWeb.HandleFunc("/api/system/build", handleGetBuildInfo)
	muxWeb.HandleFunc("/api/system/status", handleGetSystemStatus)
	muxWeb.HandleFunc("/api/login", handleAdminLogin)
	muxWeb.HandleFunc("/api/config/get", handleGetConfig)
	muxWeb.HandleFunc("/api/config/save", handleSaveConfig)
	muxWeb.HandleFunc("/api/keys/generate", handleGenerateKeys)
	muxWeb.HandleFunc("/api/keys/list", handleListKeys)
	muxWeb.HandleFunc("/api/keys/import", handleImportKeys)
	muxWeb.HandleFunc("/api/keys/delete", handleDeleteKeys)
	muxWeb.HandleFunc("/api/models/get", handleGetModelConfigs)
	muxWeb.HandleFunc("/api/models/save", handleSaveModelConfig)
	muxWeb.HandleFunc("/api/models/activate", handleActivateModel)
	muxWeb.HandleFunc("/api/rag/upload", handleRAGUpload)
	muxWeb.HandleFunc("/api/rag/status", handleRAGStatus)
	muxWeb.HandleFunc("/api/rules/get", handleGetRules)
	muxWeb.HandleFunc("/api/rules/save", handleSaveRules)
	muxWeb.HandleFunc("/api/test/chat", handleTestChat)

	webAddr := fmt.Sprintf("%s:%s", config.WebUIIP, config.WebUIPort)
	fmt.Printf("🟣 WebUI 管理后台 (HTTP) -> http://%s\n", webAddr)
	log.Fatal(http.ListenAndServe(webAddr, muxWeb))
}
