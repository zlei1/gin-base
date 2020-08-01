env = ENV['RACK_ENV'] || 'development'

require 'rubygems'
require 'bundler/setup'

Bundler.require(:default, env.to_sym)

class Application
  def self.root
    Pathname.new(File.expand_path("./", __dir__))
  end
end

require 'active_record'
include ActiveRecord::Tasks

DatabaseTasks.root = Application.root
DatabaseTasks.env = env
DatabaseTasks.database_configuration = YAML.load_file("#{Application.root}/config/database.yml")
DatabaseTasks.db_dir = 'db'
DatabaseTasks.migrations_paths = "#{Application.root}/db/migrate"
DatabaseTasks.seed_loader = Object.new
DatabaseTasks.seed_loader.instance_eval do
  def load_seed
    load "#{DatabaseTasks.db_dir}/seed.rb"
  end
end

task :environment do
  ActiveRecord::Base.configurations = DatabaseTasks.database_configuration
  ActiveRecord::Base.establish_connection DatabaseTasks.env.to_sym
end

load 'active_record/railties/databases.rake'
