<template>
  <div class="container login">
    <div class="card">
      <div class="row">
        <div class="col">
          <h2>Sign In</h2>
        </div>
      </div>
      <div class="row form--login">
        <form action="">
          <div class="col form--email">
            <div class="form-group">
              <input
                type="email"
                class="form-control"
                id="exampleInputEmail1"
                aria-describedby="emailHelp"
                placeholder="Enter email"
              />
            </div>
          </div>
          <div class="col">
            <div class="form-group">
              <input
                type="email"
                class="form-control"
                id="exampleInputEmail1"
                aria-describedby="emailHelp"
                placeholder="Enter email"
              />
            </div>
          </div>
          <div class="col btn-login">
            <button class="btn">Login</button>
          </div>
          <div class="col btn--register">
            <p>Belom Punya Akun ? <a href="">Buat Akun</a></p>
          </div>
          <div class="col btn--forgot">
            <p><a href="">Lupa Password</a></p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
.login {
  padding: 70px 0 100px 0;
  display: flex;
  justify-content: center;
}

.login .card {
  width: 577.9px;
  height: 551px;
  padding: 45px;
  box-shadow: 0px 0px 70px -18px grey;
}

.login .card h2 {
  font-weight: bold;
}

.login .card .form--login {
  width: 100%;
  margin: 39px 0 0 0;
  border-radius: 10px;
  padding: 14px 0;
  border: 2px solid grey;
}

.login .card .form--login .form--email {
  width: 100%;
  margin: 7px 0 42px 0;
}

.login .card .form--login .btn-login {
  margin: 60px 0 0 0;
}
.login .card .form--login .btn-login button {
  width: 100%;
  background-color: #e7ab3c;
  color: #fff;
  font-weight: bold;
}

.login .card .form--login .btn-login button:hover {
  opacity: 0.9;
}

.login .card .form--login .btn-login button:focus {
  outline: none;
}

.login .card .form--login .btn--register {
  margin: 10px 0 0 0;
}

.login .card .form--login .btn--register a:hover {
  color: #e7ab3c;
}
.login .card .form--login .btn--forgot {
  margin: 34px 0 0 0;
}

.login .card .form--login .btn--forgot a:hover {
  color: #e7ab3c;
}
</style>
