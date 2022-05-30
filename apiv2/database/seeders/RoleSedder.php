<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Modules\Auth\Entities\Role;

class RoleSedder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $role = new Role();
        $role->name = "superadministrator";
        $role->guard_name = "web";
        $role->save();

        $role = new Role();
        $role->name = "owner";
        $role->guard_name = "web";
        $role->save();
    }
}
