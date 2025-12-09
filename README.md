# 🩺 Health Check API em GoLang

[![Go](https://github.com/golang/go/blob/master/assets/badge.svg)](https://golang.org/)
[![Gin Gonic](https://img.shields.io/badge/Gin%20Gonic-v1.8.1-blue.svg)](https://github.com/gin-gonic/gin)
[![Prometheus](https://img.shields.io/badge/Prometheus-Client-orange.svg)](https://prometheus.io/)

Este projeto implementa uma **API RESTful de Health Check (Verificação de Saúde)** em GoLang. O objetivo é fornecer um endpoint centralizado que consulta o status de múltiplos serviços ou dependências (como bancos de dados e APIs externas) de forma concorrente, retornando o status geral e o detalhe de cada componente.

O projeto também integra o **Prometheus** para coletar métricas de latência do endpoint de saúde.

## 🚀 Tecnologias e Arquitetura

| Componente | Pacote Go | Função no Projeto |
| :--- | :--- | :--- |
| **Servidor Web** | `github.com/gin-gonic/gin` | Lida com o roteamento HTTP, expondo os endpoints `/health` e `/metrics`. |
| **Checkers** | Pacotes customizados (`checker`, `handlers`) | Implementa a lógica para verificar o status de dependências (Banco de Dados, API Externa). |
| **Métricas** | `github.com/prometheus/client_golang/prometheus` | Coleta a latência de execução do Health Check, expondo-a via endpoint `/metrics`. |
| **Context** | `context` | Usado para impor *timeouts* na execução dos verificadores, garantindo que o Health Check não demore demais. |

---

## ✨ Funcionalidades (Endpoints)

| Método HTTP | Endpoint | Descrição |
| :--- | :--- | :--- |
| **`GET`** | `/health` | **Principal Endpoint.** Executa todas as verificações de dependência em paralelo e retorna o status geral (`UP` ou `DOWN`) e o status de cada componente. |
| **`GET`** | `/metrics` | **Endpoint Prometheus.** Expõe as métricas de latência coletadas pelo *Prometheus HTTP Handler*. |

## 📦 Estrutura do Código

O projeto está organizado em módulos lógicos:

1.  **`main.go`**: Inicializa as dependências (`DatabaseChecker`, `ExternalServiceChecker`), configura o roteador Gin e define os *endpoints* `/health` e `/metrics`.
2.  **`checker/checker.go`**: Define a **interface** `DependencyChecker` e implementa as structs `DatabaseChecker` e `ExternalServiceChecker`.
    * **DatabaseChecker:** Simula uma verificação de banco de dados com um *timeout* (select com `time.After`).
    * **ExternalServiceChecker:** Simula uma verificação de API externa que falha se o tempo atual for ímpar (`time.Now().Second()%2 != 0`).
3.  **`handlers/health_handler.go`**: Contém o `HealthHandler` que executa todas as verificações fornecidas em **paralelo** (usando *goroutines* implícitas no loop de `range` + canal) e constrói a resposta JSON.
4.  **`metrics/metrics.go`**: Define e gerencia o **Gauge** do Prometheus (`HealthCheckLatency_ms`) para registrar a latência da execução do Health Check em milissegundos.

## 💾 Resposta do Endpoint `/health`

A resposta é formatada para ser legível por máquinas e geralmente inclui:

```json
{
    "status": "UP", // Status geral: UP se todos estiverem OK, DOWN se houver falha
    "timestamp": "2025-12-09T21:17:00Z",
    "components": {
        "database": "UP", // Status do checker 1
        "external_api": "DOWN" // Status do checker 2
    }
}

```
## ⚙️ Como Executar o Projeto

### 1. Pré-requisitos

Golang: Versão 1.18 ou superior.

Git: Para clonar o repositório.

### 2. Clonar e Instalar Dependências

Abra seu terminal e baixe o projeto e as dependências:

```
bash
git clone github.com/danrodsg/health-check-go.git
cd health-check-go
go mod tidy

```
### 3. Executar a API

```
bash
go run main.go
```
- O servidor estará rodando em http://localhost:8080/metrics
- O servidor estará rodando em http://localhost:8080/health
