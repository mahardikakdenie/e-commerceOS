<?php

namespace Modules\Auth\Repository;

use App\Models\User;
use Modules\Store\Entities\Store;

class AuthRepository implements AuthRepositoryInterface
{
    public function makeStore(User $userRequest, Store $store)
    {
        try {
            $user = new User();
            $user->name = $userRequest->name;
            $user->email = $userRequest->email;
            $user->password = $userRequest->password;
            $user->contact_phone = $userRequest->contact;
            $user->store_id = $userRequest->store_id;
            $user->save();
            $store = new Store();
        } catch (\Throwable $th) {
            //throw $th;
        }
    }
}
