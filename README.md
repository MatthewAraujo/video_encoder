# Video Encoder em Golang

Este é um projeto em Golang que consiste em um encoder de vídeo. A ideia principal é receber uma mensagem através do RabbitMQ, realizar o download de um vídeo armazenado em Cloud Storage, fragmentar o vídeo, convertê-lo para o formato MPEG-DASH, realizar o upload do vídeo convertido para o Cloud Storage novamente e, por fim, enviar uma notificação para a fila do RabbitMQ informando se o vídeo foi convertido com sucesso ou se houve algum erro.

## Funcionamento

1. **Recebendo mensagem do RabbitMQ**: O programa aguarda mensagens na fila do RabbitMQ. Quando uma mensagem é recebida, contendo informações sobre o vídeo a ser processado, o programa inicia o processo de conversão.

2. **Download do vídeo**: O vídeo especificado na mensagem é baixado do Cloud Storage para o sistema local.

3. **Fragmentação do vídeo**: O vídeo é fragmentado em segmentos para prepará-lo para o streaming MPEG-DASH.

4. **Conversão para MPEG-DASH**: Os segmentos do vídeo são convertidos para o formato MPEG-DASH, otimizado para streaming.

5. **Upload do vídeo convertido**: O vídeo convertido é enviado de volta ao Cloud Storage.

6. **Notificação**: Uma notificação é enviada para a fila do RabbitMQ indicando se o vídeo foi convertido com sucesso ou se houve algum erro durante o processo.

   - Se o vídeo foi convertido com sucesso, a mensagem é marcada como concluída.
   - Se ocorreu algum erro durante a conversão, a mensagem original é rejeitada e encaminhada diretamente para uma Dead Letter Exchange.

## TODO

- [ ] Implementar a conexão com o RabbitMQ para receber mensagens.
- [ ] Adicionar funcionalidade para baixar vídeos do Cloud Storage.
- [ ] Implementar a fragmentação de vídeos.
- [ ] Converter vídeos para MPEG-DASH.
- [ ] Adicionar funcionalidade para fazer o upload de vídeos convertidos para o Cloud Storage.
- [ ] Enviar notificação para o RabbitMQ indicando o status da conversão.
- [ ] Lidar com erros durante o processo de conversão.
- [ ] Configurar Dead Letter Exchange para mensagens rejeitadas.
