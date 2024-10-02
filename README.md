# Lotus-007 - Agente de Gerenciamento de Configuração

O **Lotus-007** é um agente desenvolvido para capturar e gerenciar informações de hardware e software de computadores em ambientes distribuídos. Ele é parte de uma solução completa para o gerenciamento de configurações, possibilitando o controle detalhado de ativos de TI, monitoramento contínuo e integração com sistemas de backend para armazenamento e processamento de dados.

Este projeto faz parte das atividades das disciplinas de **Projeto Temático 2** e **Gerência de Configuração**, e tem como objetivo fornecer uma plataforma robusta para coleta e envio de dados de configuração de dispositivos de forma segura e eficiente.

## Funcionalidades Principais

- **Captura de Hardware e Software**: O agente coleta informações detalhadas sobre os componentes de hardware (como processadores, discos, memória) e sobre os softwares instalados nos dispositivos monitorados.
- **Gerenciamento de Configuração**: Permite manter um inventário atualizado dos recursos de TI em toda a rede, facilitando o rastreamento e controle de mudanças.
- **Comunicação com API**: Os dados capturados pelo agente são enviados para uma API/backend central, que armazena e processa as informações de forma centralizada.
- **Segurança Integrada**: Assegura que os dados transmitidos estejam protegidos com autenticação, validação e criptografia, garantindo a integridade das informações coletadas.
- **Logs e Auditoria**: Todas as operações realizadas pelo agente, como capturas de dados e comunicações com o backend, são registradas para auditoria e monitoramento.

## Estrutura do Projeto

O **Lotus-007** é desenvolvido utilizando uma arquitetura **N-Tier** (em camadas), que garante a separação de responsabilidades e facilita a manutenção e escalabilidade do código. A seguir estão as principais camadas do sistema:

- **Processamento**: Realiza o processamento dos dados coletados do hardware e software, transformando as informações em um formato que pode ser utilizado pela camada de comunicação.
- **Data**: Gerencia a persistência local dos dados, além de lidar com a leitura e escrita de arquivos que auxiliam no funcionamento do agente.
- **Orquestração**: Coordena o fluxo de informações entre as camadas, garantindo que os dados sejam coletados, processados e enviados corretamente para o backend.
- **Segurança**: Implementa funcionalidades de segurança, incluindo autenticação e criptografia dos dados transmitidos.
- **Logging**: Registra logs de eventos e operações importantes, facilitando o monitoramento do agente e a solução de problemas.
- **Comunicação**: Gerencia a comunicação entre o agente e a API/backend, transmitindo os dados capturados de forma segura.

## Tecnologias Utilizadas

- **Linguagem**: Golang
- **Padrões de Projeto**: Factory, Singleton, Builder
- **Arquitetura**: N-Tier (em camadas)
- **Segurança**: Implementação de HMAC para validação de dados e SSL para criptografia de comunicação.

## Como Usar

1. Clone o repositório:
   git clone https://github.com/usuario/Lotus-007.git
   cd Lotus-007

2. Compile o agente:
   go build -o lotus-agent main.go

3. Execute o agente:
   ./lotus-agent

## Contribuindo

Se você deseja contribuir com o projeto, siga os passos abaixo:

1. Faça um fork do repositório.
2. Crie uma nova branch para suas alterações:
   git checkout -b feature/minha-feature

3. Envie suas alterações:
   git commit -am 'Adiciona nova feature'

4. Envie para o repositório remoto:
   git push origin feature/minha-feature

5. Abra um Pull Request.
