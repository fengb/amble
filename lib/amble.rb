#!/usr/bin/env ruby

require 'net/http'

module Amble
  autoload :VERSION, 'amble/version'

  def self.parse_headers(str)
    str.lines
    .map(&:strip)
    .grep(/[A-Za-z-]+:/)
    .map { |line| line.split(/:\s*/, 2) }
    .to_h
  end

  def self.parse_options(header_file, path_file)
    {
      paths: IO.readlines(path_file).map(&:strip),
      headers: parse_headers(IO.read(header_file))
    }
  end

  def self.run(headers:, paths:)
    paths.each do |path|
      uri = URI("http://#{headers['Host']}#{path}")
      req = Net::HTTP::Get.new(uri)

      headers.each do |key, value|
        req[key] = value
      end

      Net::HTTP.start(uri.host, uri.port) do |http|
        res = http.request req
        filename = path.sub(/^\//, '').tr('/', '-')
        puts "####{filename}###"
        puts "#{res.body}\n"
      end
    end
  end
end
