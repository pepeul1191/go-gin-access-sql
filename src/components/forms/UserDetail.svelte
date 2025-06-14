<script>
  import { onMount } from 'svelte';

  export const clean = () => {

  }

  export let user = {
    id: null,
    username: '',
    password: '',
    activation_key: '',
    reset_key: '',
    email: '',
    activated: false,
    created: '',
    updated: ''
  };

  let btnDisabled = false;

  // Funciones auxiliares
  function generateKey(length = 12) {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    return Array.from({ length }, () => chars[Math.floor(Math.random() * chars.length)]).join('');
  }

  function regeneratePassword() {
    user.password = generateKey(16); // Cambiar si usas un backend que genera el hash
  }

  function regenerateActivationKey() {
    user.activation_key = generateKey(10);
  }

  function regenerateResetKey() {
    user.reset_key = generateKey(10);
  }

  function submitForm() {
    // Aquí haces el fetch/axios PUT a tu API
    console.log('Enviar usuario actualizado:', user);
  }
</script>

<style>
  h4{
    font-size: 1.075rem;
  }
</style>

<div class="row subtitle-row">
  <h4 class="subtitle">Datos Generales</h4>
</div>
<div class="row align-items-end">
  <div class="col-4">
    <label class="form-label">Usuario</label>
    <input type="text" class="form-control" bind:value={user.username}>
  </div>

  <div class="col-5">
    <label class="form-label">Correo</label>
    <input type="email" class="form-control" bind:value={user.email}>
  </div>

  <div class="col-3">
    <button class="btn btn-success">
      <i class="fa fa-user"></i> Crear
    </button>
  </div>
</div>

<div class="col-12 mt-2">
  <div class="mb-3">
    <label class="form-label">Contraseña</label>
    <div class="input-group">
      <input type="text" class="form-control" bind:value={user.password}>
      <button type="button" class="btn btn-outline-secondary" on:click={regeneratePassword}><i class="fa fa-random"></i> Regenerar</button>
      <button type="button" class="btn btn-outline-secondary" on:click={regeneratePassword}><i class="fa fa-random"></i> Copiar</button>
      <button type="button" class="btn btn-outline-secondary" on:click={regeneratePassword}><i class="fa fa-random"></i> Guardar</button>
    </div>
  </div>
</div>

<div class="row g-3 mb-1">
  <!-- Columna 1: Activar Cuenta -->
  <div class="col-md-4">
    <h4 class="subtitle mb-3">Estado de Cuenta</h4>
    <div class="d-flex justify-content-start">
      <button class="btn btn-primary">
        <i class="fa fa-check"></i> Activar Cuenta
      </button>
    </div>
  </div>

  <!-- Columna 2: Enviar Solicitudes -->
  <div class="col-md-8">
    <h4 class="subtitle mb-3">Enviar Solicitudes a Correo</h4>
    <div class="d-flex flex-wrap gap-2">
      <button class="btn btn-info">
        <i class="fa fa-envelope"></i> Activación de Cuenta
      </button>
      <button class="btn btn-warning">
        <i class="fa fa-refresh"></i> Cambio de Contraseña
      </button>
    </div>
  </div>
</div>
