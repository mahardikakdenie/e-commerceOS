<?php

use Illuminate\Http\Request;
use Modules\Auth\Http\Controllers\AuthController;
use Modules\Auth\Http\Controllers\RegisterController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/auth', function (Request $request) {
    return $request->user();
});

Route::prefix("/auth")->group(function () {
    Route::post("/login/owner", [AuthController::class, "loginOwner"]);
    Route::post("/login/admin", [AuthController::class, "loginAdmin"]);
    Route::post("/register/admin", [RegisterController::class, "registerAsSuperAdministrator"]);
    Route::post("/register", [RegisterController::class, 'makeStore']);
});
