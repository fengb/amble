# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'amble/version'

Gem::Specification.new do |spec|
  spec.name          = 'amble'
  spec.version       = Amble::VERSION
  spec.authors       = ['Benjamin Feng']
  spec.email         = ['contact@fengb.me']

  spec.summary       = 'Quickly dumps a bunch of requests for comparison.'
  spec.description   = 'Quickly dumps a bunch of requests for comparison.'
  spec.homepage      = 'https://github.com/fengb/amble'
  spec.license       = 'MIT'

  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.bindir        = 'exe'
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ['lib']

  spec.add_development_dependency 'bundler', '~> 1.14'
  spec.add_development_dependency 'rake', '~> 10.0'
  spec.add_development_dependency 'rspec', '~> 3.0'
end
