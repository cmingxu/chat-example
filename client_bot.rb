require 'socket'
require 'json'

tcp = TCPSocket.new "localhost", 10000
message = {"FromId"=> 10, "ToId"=> 12, "ChannelId"=> 1, "MessageType"=> 1, "MessageContent"=> "Hello World"}
read_thread = Thread.new do
  while true do
    puts "xxxxxxxxxx"
    mes = tcp.gets
    puts "Client Receiving .."
    puts JSON.parse(mes)
  end
end

write_thread = Thread.new do
  while true
    $stdout.sync = true
    message["MessageContent"] = "xxx#{rand(100)}"
    puts "Client Writing: #{message.to_json}"
    $stdout.flush
    tcp.write message.to_json
    tcp.write "\n"
    sleep 3
  end
end

read_thread.join
write_thread.join
