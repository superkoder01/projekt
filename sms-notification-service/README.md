# SMS Notification Service

### RabbitMQ Properties
* Queue name: "sms-que"
* Message content type: "application/json"
* Message Id (Optional): UUID


### SMS RabbitMQ Message Model
```
{
    "sender": "00123123123" //Gateway API Sender
    "msisdn": ["0048111222333", "111222333"]
    "text": "SMSTEXT"
}
```