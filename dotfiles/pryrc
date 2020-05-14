# Add the following to each project Gemfile:
#   gem 'pry-rails', :group => :development

# References:
# - https://github.com/pry/pry/wiki/Pry-rc
# - https://github.com/rweng/pry-rails

begin
  require 'awesome_print'
  Pry.config.print = proc { |output, value| Pry::Helpers::BaseHelpers.stagger_output("=> #{value.ai}", output) }
rescue LoadError => err
  puts "Couldn't load awesome_print "
end

if defined?(PryRails::RAILS_PROMPT)
  Pry.config.prompt = PryRails::RAILS_PROMPT
end
