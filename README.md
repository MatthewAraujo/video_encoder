# Video Encoder in Golang

This is a Golang project that consists of a video encoder. The main idea is to receive a message via RabbitMQ, download a video stored in Cloud Storage, fragment the video, convert it to MPEG-DASH format, upload the converted video back to Cloud Storage, and finally, send a notification to the RabbitMQ queue informing whether the video was successfully converted or if there was any error.

## Operation

1. **Receiving message from RabbitMQ**: The program waits for messages in the RabbitMQ queue. When a message is received, containing information about the video to be processed, the program initiates the conversion process.

2. **Video download**: The video specified in the message is downloaded from Cloud Storage to the local system.

3. **Video fragmentation**: The video is fragmented into segments to prepare it for MPEG-DASH streaming.

4. **Conversion to MPEG-DASH**: The video segments are converted to MPEG-DASH format, optimized for streaming.

5. **Upload of converted video**: The converted video is sent back to Cloud Storage.

6. **Notification**: A notification is sent to the RabbitMQ queue indicating whether the video was successfully converted or if there was any error during the process.

   - If the video was successfully converted, the message is marked as completed.
   - If an error occurred during conversion, the original message is rejected and forwarded directly to a Dead Letter Exchange.

## TODO

- [ ] Implement RabbitMQ connection to receive messages.
- [ ] Add functionality to download videos from Cloud Storage.
- [ ] Implement video fragmentation.
- [ ] Convert videos to MPEG-DASH.
- [ ] Add functionality to upload converted videos to Cloud Storage.
- [ ] Send notification to RabbitMQ indicating conversion status.
- [ ] Handle errors during the conversion process.
- [ ] Configure Dead Letter Exchange for rejected messages.
