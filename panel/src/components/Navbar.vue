<template>
  <nav class="navbar navbar-expand-xl navbar-light bg-grey">
    <div class="container">
      <h3 class="text-black cursor-pointer text--brand font-bold" href="#">
        E-Commerce
      </h3>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div
        class="collapse navbar-collapse navbar--content"
        id="navbarSupportedContent"
      >
        <ul class="ms-auto navbar-nav">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="#">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Dashboard</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Chart</a>
          </li>
          <li class="nav-item">
            <button class="btn btn-secondary btn--login--width" type="submit">
              login
            </button>
          </li>
        </ul>
      </div>
      <div></div>
    </div>
  </nav>
</template>

<script>
export default {};
</script>

<style></style>
