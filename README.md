# gonotify
Notification Service that sends push notification to websocket clients connected to it. 

This is deployed at gonotify.herokuapp.com

Websocket clients can connect to it wss://gonotify.herokuapp.com/ws/\<clientId\>   

Replace client Id with actual identiifer used to uniquely identify the client , such as userId. Note that same client can connect from multiple end points, so this is not a end point Id , but a client id.

For example, say a client awesome can connect to wss://gonotify.herokuapp.com/ws/awesome  from multiple browsers / apps. 

Any applciation that has to send message to all connected end points of a client, can invoke the API (Http POST) https://gonotify.herokuapp.com/notify/\<clientId\>  and the message will be delivered to all connected endpoints for the specified client.

For example, to send the message "Hello Friend!" to clients connected using clientId awesome, the API URL and payload has to be:

  * URL: https://gonotify.herokuapp.com/notify/awesome
  * Http method: POST
  * Payload:  {"message":"Hello Friend!"}
  
The value of message in the payload can be another JSON, but enclosed with in double quotes (as a string).
 * Example: 
   * { "message": "{"msg":"Hello from User1"}"} 

The API response has a status field, which if true indicates that there were one or more client endpoints connected to those the message is delivered. If the status is false, it indicates there are no endpoints connected using the specified client id and message is dropped immediately. 
Note that in both these cases, the http status code will be 200. 

Coming Soon:

The component that gonotify uses, gotell, supports sending APNS push notification to devices.  However this is limited to one application instance per application that has to send push notification. Once this is enhanced, it will be made part of gonotify.

