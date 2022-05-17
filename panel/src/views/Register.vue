<template>
  <div class="container register">
    <div class="card">
      <div class="row">
        <div class="col">
          <h2>Sign Up</h2>
        </div>
      </div>
      <form class="row g-3">
        <div class="col-md-12">
          <label for="inputEmail4" class="form-label">Email</label>
          <input type="email" class="form-control" id="inputEmail4" />
        </div>
        <div class="col-md-6">
          <label for="inputEmail4" class="form-label">Email</label>
          <input type="email" class="form-control" id="inputEmail4" />
        </div>
        <div class="col-md-6">
          <label for="inputPassword4" class="form-label">Password</label>
          <input type="password" class="form-control" id="inputPassword4" />
        </div>

        <div class="col-12">
          <label for="inputAddress2" class="form-label">Contact</label>
          <input
            type="text"
            class="form-control"
            id="inputAddress2"
            placeholder="Contact"
          />
        </div>

        <!-- <div class="col-12">
          <div class="form-check">
            <input class="form-check-input" type="checkbox" id="gridCheck" />
            <label class="form-check-label" for="gridCheck">
              Check me out
            </label>
          </div>
        </div> -->
        <div class="col-12">
          <button type="submit" class="btn btn-primary">Sign in</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
.register {
  padding: 70px 0 100px 0;
  display: flex;
  justify-content: center;
}

.register .card {
  width: 577.9px;
  height: 100%;
  padding: 45px;
  box-shadow: 0px 0px 70px -18px grey;
}

.register .btn {
  width: 100%;
  height: 50px;
  border-radius: 0;
  border: none;
  color: white;
  font-size: 18px;
  font-weight: bold;
  margin-top: 20px;
}
</style>
