# Lotus-007 - Branch Develop

Repositório para o agente do projeto Lotus, utilizado nas aulas de Projeto Temático 2 e Gerência de Configuração.

---

## Visão Geral

O **Lotus-007** é um agente desenvolvido para capturar informações de hardware e software de computadores, sendo parte fundamental do gerenciamento de configurações em ambientes distribuídos.

Esta branch está focada em converter o código inicial (Preview) para um design orientado a objetos, aplicando padrões de projeto (Design Patterns) e implementando uma arquitetura N-Tier com as seguintes camadas:

- **Processamento**: Responsável por processar os dados capturados do hardware e software.
- **Data**: Camada responsável pelo gerenciamento e persistência dos dados.
- **Orquestração**: Coordenadora entre as diferentes camadas do sistema, garantindo o fluxo correto entre captura, processamento e envio de dados.
- **Segurança**: Cuida dos aspectos relacionados à segurança, como autenticação, validação e criptografia dos dados transmitidos.
- **Logging**: Responsável pela captura e armazenamento de logs de eventos e operações do agente.
- **Comunicação**: Gerencia a comunicação entre o agente e a API/backend, enviando os dados capturados e processados.

---

## Objetivos da Branch

Nesta branch, o foco é a refatoração do código para:

- Implementar um design orientado a objetos (OO).
- Adotar padrões de projeto (Design Patterns) como Factory, Singleton e Builder para modularizar e organizar melhor o código.
- Melhorar a arquitetura, dividindo o sistema em camadas bem definidas (N-Tier).
- Garantir que cada camada esteja bem desacoplada e tenha responsabilidades claras.

---

## Padrões de Projeto Utilizados

- **Factory**: Para criar diferentes implementações de captura de hardware e software, além de ser utilizado também para o logging.
- **Singleton**: Para garantir que certas instâncias, como a de Logging e Segurança, existam apenas uma vez durante a execução do programa.
- **Builder**: Para estruturar a orquestração na formação dos JSONs capturados.
  
---

## Como Contribuir

1. Faça um fork do repositório.
2. Crie uma nova branch para suas alterações: `git checkout -b feature/minha-feature`
3. Envie suas alterações: `git commit -am 'Adiciona nova feature'`
4. Envie para o repositório remoto: `git push origin feature/minha-feature`
5. Abra um Pull Request.

---
