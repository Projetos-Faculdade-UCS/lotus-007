# Lotus-007

Repositório para o agente do projeto Lotus da aula de Projeto Temático 2 e também utilizado na aula de Gerência de Configuração.

---

## Visão Geral

O **Lotus-007** é um agente desenvolvido para capturar informações de hardware e software de computadores, sendo parte fundamental do gerenciamento de configurações em ambientes distribuídos. 

O projeto adota o modelo N-Tier e está sendo desenvolvido em etapas, com foco em experimentar diferentes abordagens e capturar a melhor forma de coletar dados.

---

## Preview

Esta branch contém o código feito apenas para testar conceitos de como o agente poderia funcionar, sem preocupação com implementação final, arquitetura robusta ou orientação a objetos.

Com este Preview, foi possível:
- Testar diferentes métodos de captura de dados de hardware.
- Identificar o melhor método para fazer as capturas.
- Definir a estrutura básica das camadas que serão utilizadas no modelo N-Tier.

Essa branch servirá de base para o lançamento de um Alpha em breve.

---

## Compatibilidade

Esta versão do agente foi testada e é compatível com:

- **Windows 10**
- **Windows 11**

### Requisitos Especiais

- **Desativar o antivírus**: Esta versão do agente requer que o antivírus esteja desativado, pois ela executa operações que envolvem alterações no registro do Windows e no PowerShell, o que pode ser bloqueado por ferramentas de segurança.

---

## Dependências

Este projeto requer a seguinte dependência para funcionar corretamente:

- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) v0.24.0

Você pode instalá-la com o comando:

```bash
go get golang.org/x/sys@v0.24.0
```

## Instruções de Execução

Para executar o agente, siga os passos abaixo:

1. Navegue até a pasta onde o arquivo `main.go` está localizado.

2. Compile o agente utilizando o comando:

    ```bash
    go build
    ```

3. Execute o agente utilizando o seguinte comando:

    ```bash
    .\agent.exe
    ```

Certifique-se de estar na pasta correta onde o arquivo `main.go` está antes de executar os comandos.

**Nota**: Pode ser necessário inicializar o módulo Go caso ele não esteja configurado, utilizando o seguinte comando:

```bash
go mod init

