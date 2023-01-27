# Programando um echo server

Quando falamos de criar um programa cliente e servidor, umas das primeiras aplicações que vem à cabeça é um _echo server_. Esta aplicação permite que um cliente mande uma mensagem ao servidor e este servidor responda com a mesma mensagem enviada - por isso _eco_.

Como essa relação cliente e servidor impacta nosso dia-a-dia? O que é um servidor? O que é um cliente?

## Servidor

Um servidor é um programa de computador que providencia e recebe informações de clientes, ou outros servidores. A internet é cheia de servidores, todos eles provendo soluções distintas. Para que um cliente, por exemplo um navegador, possa conectar em um servidor, o cliente precisa saber o IP e a porta na qual aquele servidor está escutando. Nem todo servidor precisa estar em uma rede pública. Eles podem estar em redes privadas ou até mesmo na sua própria rede local, usando como endereço o famoso _localhost_. IP, porta e protocolo de comunicação são informações essenciais para que um cliente ou servidor precisam para se comunicar com servidores.

## Cliente

Clientes são ferramentas usadas para conectar e trocar informações com os serviços oferecidos por servidores. Um cliente pode ser tanto um aplicativo, um navegador, ou a ferramenta _curl_. Ambas ferramentas sabem como conectar em uma página web, por exemplo, seguindo o protocolo HTTP.
