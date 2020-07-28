require_relative 'application'
require 'active_support/core_ext/string/strip'

namespace :db do
  # rake db:generate_migration[create_tablename]
  # rake db:generate_migration\[create_tablename\] if zsh: no matches found ...
  desc "generate a migration"
  task :generate_migration, [:name] do |_, args|
    name, version = args[:name], Time.now.utc.strftime("%Y%m%d%H%M%S")

    ActiveRecord::Migrator.migrations_paths.each do |directory|
      next unless File.exist?(directory)
      migration_files = Pathname(directory).children
      if duplicate = migration_files.find { |path| path.basename.to_s.include?(name) }
        abort "Another migration is already named \"#{name}\": #{duplicate}."
      end
    end

    filename = "#{version}_#{name}.rb"
    dirname  = ActiveRecord::Migrator.migrations_paths.first
    path     = File.join(dirname, filename)
    ar_maj   = ActiveRecord::VERSION::MAJOR
    ar_min   = ActiveRecord::VERSION::MINOR
    base     = "ActiveRecord::Migration"
    base    += "[#{ar_maj}.#{ar_min}]" if ar_maj >= 5

    FileUtils.mkdir_p(dirname)
    File.write path, <<-MIGRATION.strip_heredoc
      class #{name.camelize} < #{base}
        def change
        end
      end
    MIGRATION

    puts path
  end
end

