<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth as FacadesAuth;

class AuthController extends Controller
    {
        public function tampilRegistrasi(){
            return view('register');
        }

        public function submitRegistrasi(Request $request){
            // Run validation first
            $request->validate([
                'name' => 'required|string|max:255|unique:users,name',
                'email' => 'required|email|max:255|unique:users,email',
                'password' => 'required|min:6|',
            ]);

            // Simpan user baru
            User::create([
                'name' => $request->name,
                'email' => $request->email,
                'password' => bcrypt($request->password), // Enkripsi password
            ]);

            return redirect()->route('login.tampil')->with('success', 'Registrasi berhasil! Silakan login.');
        }

        public function tampilLogin(){
            return view('login');
        }

        public function submitLogin(Request $request){
            $data = $request->only('email', 'password');
            if (FacadesAuth::attempt($data)){
                $request->session()->regenerate();
                return redirect()->route('home');
            } else {
                return redirect()->back()->with('error', 'Login gagal, silahkan coba lagi');
            }
        }

        public function submitLogout(){
            FacadesAuth::logout();
            return redirect()->route('login.tampil');
        }
    };
