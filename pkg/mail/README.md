# Pacote de Email
Para criar e enviar um email usando o pacote de email, você vai precisar dessas três coisas:

- Criar a Configuração de Email: Você precisa configurar o email com suas credenciais (email, senha, provedor, host e porta).

- Compor a Mensagem de Email: Defina o destinatário, o conteúdo da mensagem, o assunto e qualquer anexo.

- Enviar o Email: Use a configuração e a mensagem para enviar o email.

e utilizar o pacote:
```
package mail
```

## 1. Criar configuração de email
Para criar um email, você deve primeiro configurar o email. Para isso, use a função **CreateConfig**, passando o email, a senha (secret), o provedor, o endereço do host e a porta do host.

assim:

cfg, err := CreateConfig("example@gmail.com", "secret", "smtp.gmail.com", "smtp.gmail.com", "587")

e vai retornar um struct como essa:

```
{
    EmailAdress         string
    EmailSecretPassword string
    Provider            string
    HostAddress         string
    HostPort            string
}
```

## 2. Criar a Instância do Email
Para criar o próprio email, passe a configuração criada na etapa anterior: 

mail := NewMail(cfg)

Isso retornará uma struct com a configuração passada.

assim: 

```
{
    Config *Config
}
```

## 3. Criar a Mensagem de Email
Para enviar um email, é necessário definir o conteúdo do email. Use a função **NewMessage** para criar uma nova mensagem. Passe o(s) destinatário(s) e o conteúdo HTML:

message := NewMessage("recipient@example.com", "<h1>Hello, World!</h1>")

Isso retorna uma struct Mensagem com os seguintes campos:

```
{
    From        string
    To          []string
    Html        string
    Subject     string
    Cc          []string
    Bcc         []string
    ReplyTo     []string
    Attachments Attachment
}
```

## 4. Adicionando Anexos
Se o email precisar incluir anexos, use o método AttachFile:

err := message.AttachFile("/path/to/file.txt")

## Enviar o Email
Para enviar o email composto, use o método Enviar da struct Email:

err := mail.Send(message)

## 6. Exemplo Completo
Aqui está um exemplo completo de como configurar a configuração, criar o email, anexar um arquivo e enviá-lo:

```
cfg, err := CreateConfig("example@gmail.com", "secret", "smtp.gmail.com", "smtp.gmail.com", "587")
if err != nil {
    log.Fatalf("Failed to create email config: %v", err)
}

mail := NewMail(cfg)

message := NewMessage("recipient@example.com", "<h1>Hello, World!</h1>")
message.Subject = "Welcome Email"

err = message.AttachFile("/path/to/file.txt")
if err != nil {
    log.Fatalf("Failed to attach file: %v", err)
}

err = mail.Send(message)
if err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```