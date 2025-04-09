# Mini Kubernetes em Go

Este projeto é uma simulação de um orquestrador de containers fake, inspirado no Kubernetes, escrito em Go. Ele utiliza conceitos de sistemas distribuídos, goroutines e gRPC para simular agendamentos e comunicação entre nós.

## Funcionalidades

- Simulação de containers como goroutines
- Comunicação entre nós com gRPC (estrutura inicial)
- Controle de status e logs dos containers
- Implementação básica do algoritmo de consenso Raft (em progresso)

## Execução

```bash
go run main.go
