# ruby 1.8.7 compatible
# Reference: http://www.rakeroutes.com/blog/customize-your-irb/

require 'rubygems'
require 'irb/completion'

begin
  require "awesome_print"
  AwesomePrint.irb!
rescue LoadError => err
  warn "Couldn't load awesome_print: #{err}"
end

# irb history
IRB.conf[:EVAL_HISTORY] = 1000
IRB.conf[:SAVE_HISTORY] = 1000
IRB.conf[:HISTORY_FILE] = File::expand_path("~/.irbhistory")

# load .railsrc in rails environments
railsrc_path = File.expand_path('~/.irbrc_rails')
if ( ENV['RAILS_ENV'] || defined? Rails ) && File.exist?( railsrc_path )
  begin
    load railsrc_path
  rescue Exception
    warn "Could not load: #{ railsrc_path } because of #{$!.message}"
  end
end

class Object
  def interesting_methods
    case self.class
    when Class
      self.public_methods.sort - Object.public_methods
    when Module
      self.public_methods.sort - Module.public_methods
    else
      self.public_methods.sort - Object.new.public_methods
    end
  end
end

module Kernel
  def require_relative(file)
    $:.unshift Dir.pwd
    require file
  end

  def guid(s)
    s.scan(/[a-f0-9-]{36}/).first
  end
end

#   make sure I'm always the first user in the database so that this works
def me
  User.first
end

#   quick reload the IRB session
def r
  reload!
end

#   quick reload FactoryBot
def fr
  FactoryBot.reload
end
