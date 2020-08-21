# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `rails
# db:schema:load`. When creating a new database, `rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2020_07_29_023820) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "admins", comment: "管理员", force: :cascade do |t|
    t.string "name", null: false, comment: "姓名"
    t.string "phone", null: false, comment: "手机号"
    t.string "encrypted_password", null: false, comment: "密码"
    t.datetime "current_sign_in_at", comment: "当前登入时间"
    t.datetime "last_sign_in_at", comment: "上次登入时间"
    t.string "current_sign_in_ip", comment: "当前登入Ip"
    t.string "last_sign_in_ip", comment: "上次登入Ip"
    t.integer "status", default: 0, comment: "状态"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["phone"], name: "index_admins_on_phone", unique: true
  end

  create_table "users", comment: "用户", force: :cascade do |t|
    t.string "name", null: false, comment: "姓名"
    t.string "phone", null: false, comment: "手机号"
    t.string "encrypted_password", null: false, comment: "密码"
    t.datetime "current_sign_in_at", comment: "当前登入时间"
    t.datetime "last_sign_in_at", comment: "上次登入时间"
    t.string "current_sign_in_ip", comment: "当前登入Ip"
    t.string "last_sign_in_ip", comment: "上次登入Ip"
    t.integer "status", default: 0, comment: "状态"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["phone"], name: "index_users_on_phone", unique: true
  end

end
