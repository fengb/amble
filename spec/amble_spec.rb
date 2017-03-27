require 'spec_helper'

RSpec.describe Amble do
  describe '.parse_headers' do
    it 'parses simple headers' do
      headers = Amble.parse_headers('Host: example.com')
      expect(headers).to eq('Host' => 'example.com')
    end

    it 'parses stupid headers' do
      headers = Amble.parse_headers('User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3005.0 Safari/537.36')
      expect(headers).to eq('User-Agent' => 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3005.0 Safari/537.36')
    end

    it 'ignores non-headers' do
      headers = Amble.parse_headers('GET /api HTTP/1.1')
      expect(headers).to eq({})
    end

    it 'parses multiple headers' do
      headers = Amble.parse_headers <<~END
                  Host: example.com
                  X-Requested-With: XMLHttpRequest
                END
      expect(headers).to eq(
        'Host' => 'example.com',
        'X-Requested-With' => 'XMLHttpRequest',
      )
    end
  end
end
