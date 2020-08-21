class CreateAdmins < ActiveRecord::Migration[6.0]
  def change
    create_table :admins, comment: "管理员" do |t|
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

    add_index :admins, :phone, unique: true
  end
end
