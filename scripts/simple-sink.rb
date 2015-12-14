require 'sinatra'

post '*' do
  puts params.inspect
end
