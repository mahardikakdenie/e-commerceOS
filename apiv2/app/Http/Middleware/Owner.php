<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Brryfrmnn\Transformers\Json;

class Owner
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure(\Illuminate\Http\Request): (\Illuminate\Http\Response|\Illuminate\Http\RedirectResponse)  $next
     * @return \Illuminate\Http\Response|\Illuminate\Http\RedirectResponse
     */
    public function handle(Request $request, Closure $next, $role)
    {
        if ($request->user()->hasRole($role)) {
            return $next($request);
        }
        return Json::exception('You are not authorized to access this page.');
    }
}
