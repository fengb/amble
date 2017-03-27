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

  describe '.parse_options' do
    describe '[:headers]' do
      after do
        $stdin = STDIN
      end

      it 'defaults to {}' do
        opts = Amble.parse_options([])
        expect(opts[:headers]).to eq({})
      end

      it 'parses_headers(stdin) if not tty' do
        $stdin = StringIO.new('this is sparta')
        expect(Amble).to receive(:parse_headers).with('this is sparta').and_return(123)
        opts = Amble.parse_options([])
        expect(opts[:headers]).to eq(123)
      end

      it 'parses_headers(file) via -f "file"' do
        expect(IO).to receive(:read).with('foobar').and_return('this is sparta')
        expect(Amble).to receive(:parse_headers).with('this is sparta').and_return('dumdum')
        opts = Amble.parse_options(['-f', 'foobar'])
        expect(opts[:headers]).to eq('dumdum')
      end
    end

    describe '[:paths]' do
      it 'defaults to {}' do
        opts = Amble.parse_options([])
        expect(opts[:paths]).to eq([])
      end

      it 'converts regular arguments' do
        opts = Amble.parse_options(['/api', '/google'])
        expect(opts[:paths]).to eq(['/api', '/google'])
      end
    end
  end
end
