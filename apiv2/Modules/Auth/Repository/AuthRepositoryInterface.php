<?php

namespace Modules\Auth\Repository;

use App\Models\User;
use Illuminate\Support\Facades\Request;
use Modules\Store\Entities\Store;

class AuthRepositoryInterface
{
    public function makeStore(User $user, Store $store);
}
