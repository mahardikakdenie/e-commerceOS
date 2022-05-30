<?php

use Illuminate\Http\Request;
use Modules\Store\Http\Controllers\StoreController;

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

Route::middleware('auth:api')->get('/store', function (Request $request) {
    return $request->user();
});

Route::prefix('store')->middleware(['auth:sanctum', "owner:owner"])->group(function () {
    Route::get('', [StoreController::class, 'index']);
    // Route::get('{id}', [StoreController::class, 'show']);
});
