require 'socket'
require 'json'

tcp = TCPSocket.new "localhost", 10000
message = {"FromId"=> 10, "ToId"=> 12, "ChannelId"=> 1, "MessageType"=> 1, "MessageContent"=> "Hello World"}
#Thread.new do
  #while true do
    #mes = tcp.read
    #puts JSON.parse(mes)
  #end
#end.join

while true
  $stdout.sync = true
  message["MessageContent"] = "xxx#{rand(100)}"
  puts "Writing: #{message.to_json}"
  $stdout.flush
  tcp.write message.to_json
  tcp.write "\n"
  sleep 1
end
