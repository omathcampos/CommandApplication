# 📦 Aplicação de Linha de Comando para Busca de IPs e Nomes de Servidores

# 📜 Descrição
# Esta aplicação de linha de comando permite buscar endereços IP 🌐 e nomes de servidores 🖥 da internet.
# Usando a biblioteca `urfave/cli`, a aplicação é fácil de usar e fornece uma interface simples para diversos comandos.

# ⚙️ Funcionalidades
# - 🔍 **Busca de endereços IP** a partir de um nome de host.
# - 🔎 **Busca de nomes de servidores** a partir de um nome de host.

# 📋 Pré-requisitos
# - Go instalado no seu sistema. É recomendado usar a versão 1.11 ou superior.
# - Dependências do projeto devem ser instaladas (veja instruções abaixo).

# 🚀 Instalação

# 1. **Clone este repositório:**
git clone git@github.com:omathcampos/CommandApplication.git
cd CommandApplication

# 2. **Instalar dependências:** Você pode usar o seguinte comando para baixar as dependências do projeto:
go get

# 💻 Como Usar

# 👩‍💻 Execução Direta
# Para executar a aplicação diretamente a partir do código fonte:

# 1. Navegue até o diretório onde a aplicação está localizada.

# 2. Use os seguintes comandos para buscar por servidores ou endereços IP:

# - **Buscar por Servidores:**
go run main.go Servidores --host <url>
# **Exemplo:**
go run main.go Servidores --host amazon.com.br

# - **Buscar por IP:**
go run main.go IP --host <nome do site>
# **Exemplo:**
go run main.go IP --host google.com

# ### 🛠 Compilação
# Para compilar a aplicação em um executável:
go build -o linha-de-comando

# Após a compilação, você pode executar o executável:

# - **Buscar por Servidores:**
./linha-de-comando Servidores --host <url>

# - **Buscar por IP:**
./linha-de-comando IP --host <nome do site>

# 📌 Exemplos de Uso
# - Buscar endereços IP para um host:
./linha-de-comando IP --host amazon.com.br

# - Buscar servidores para um host:
./linha-de-comando Servidores --host google.com