require 'json'
require 'digest'
require './data.rb'

currency_hash = {}

JSON.parse(DATA).map do |c|
  currency_hash[Digest::SHA256.hexdigest c['name']] = c
end

puts JSON.generate(currency_hash)