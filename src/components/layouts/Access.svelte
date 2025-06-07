<script>
  import { onMount } from 'svelte';

  let error = null;
  let usernameInput;

  onMount(() => {
    // Obtener los parámetros de la URL al cargar el componente
    const urlParams = new URLSearchParams(window.location.search);
    error = urlParams.get('error') || '';
    usernameInput.focus();
    setTimeout(() => {
      error = null;
    }, 4300);
  });
</script>
  
<style>

</style>

<div class="container d-flex justify-content-center align-items-center" style="min-height: 100vh;">
  <div class="card shadow p-4" style="width: 100%; max-width: 400px;">
    <h3 class="text-center mb-4">Iniciar Sesión</h3>
    {#if error}
      <div class="alert alert-danger" role="alert">
        {error}
      </div>
    {/if}
    <form action="/login" method="POST">
      <div class="mb-3">
        <label for="username" class="form-label">Usuario</label>
        <input type="text" class="form-control" id="username" name="username" required bind:this={usernameInput}>
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Contraseña</label>
        <input type="password" class="form-control" id="password" name="password" required>
      </div>
      <button type="submit" class="btn btn-primary w-100">Ingresar</button>
    </form>
  </div>
</div>