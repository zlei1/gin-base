class CreateUsers < ActiveRecord::Migration[6.0]
  def change
    create_table :users, comment: "用户" do |t|
      t.string :code, null: false, comment: "编号"
      t.string :name, null: false, comment: "姓名"
      t.string :phone, null: false, comment: "手机号"
      t.string :encrypted_password, null: false, comment: "密码"
      t.datetime :current_sign_in_at, comment: "当前登入时间"
      t.datetime :last_sign_in_at, comment: "上次登入时间"
      t.string :current_sign_in_ip, comment: "当前登入Ip"
      t.string :last_sign_in_ip, comment: "上次登入Ip"
      t.integer :status, default: 0, comment: "状态"
      t.timestamps
    end

    add_index :users, :code, unique: true
    add_index :users, :phone, unique: true
  end
end
