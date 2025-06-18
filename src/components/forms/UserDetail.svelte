<script>
  import { onMount } from 'svelte';
  import axios from 'axios';
  import { createEventDispatcher } from 'svelte';

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

  let message = {
    text: '',
    status: ''
  };

  let messagePassword = ''
  let btnsDisabled = true;

  const dispatch = createEventDispatcher();

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

  const cleanMessage = (dispatchToParent) => {
    setTimeout(() => {
      message = {
        text: '',
        status: ''
      };
      btnDisabled = false;
      if (dispatchToParent){
        dispatch('saved', { id: user.id });
      }
    }, 4300);
  }

  const handleSave = async (event) => {
    event.preventDefault(); 
    // Crea un objeto con los datos del formulario
    try {
      // Enviar datos usando Axios
      let response;
      btnDisabled = true;
      if(user.id == null){
        const formData = {
          username: user.username,
          email: user.email
        };
        response = await axios.post(BASE_URL + 'apis/v1/users', formData, {
          headers: {
            'Content-Type': 'application/json',
          },
        });
        user.id = response.data.id;
        message.text = 'Se ha creado el sistema';
        message.status = 'success';
        // notificar al padre que se ha actualizado algo
        btnsDisabled = false;
        cleanMessage(true);
      }else{
        response = await axios.put(BASE_URL + 'apis/v1/systems', formData, {
          headers: {
            'Content-Type': 'application/json',
          },
        });  
        updated = toDatetimeLocalWithSeconds(response.data.updated);
        message.text = 'Se ha editado el usuario';
        message.status = 'success';
        cleanMessage(true);
      }
      console.log('Datos enviados con éxito:', response.data);
      // Puedes manejar la respuesta aquí, por ejemplo, mostrar un mensaje de éxito
    } catch (error) {
      console.error('Error al enviar los datos:', error);
      message.text = 'Error al grabar el usuario';
      message.status = 'danger';
      cleanMessage(false);
      // Maneja el error (puedes mostrar un mensaje de error en la interfaz de usuario)
    }
  }

  const savePassword = async (event) => {
    event.preventDefault(); 
    // Crea un objeto con los datos del formulario
    try {
      // Enviar datos usando Axios
      let response;
      btnDisabled = true;
      if(user.id != null){
        const formData = {
          password: user.password,
        };
        response = await axios.put(BASE_URL + 'apis/v1/users/' + user.id + '/password', formData, {
          headers: {
            'Content-Type': 'application/json',
          },
        });
        message.text = 'Se ha actualizado la contraseña del usuario';
        message.status = 'success';
        // notificar al padre que se ha actualizado algo
        cleanMessage(false);
      }else{
        message.text = 'No puede actualizar contraseña sin crear primero al usuario';
        message.status = 'danger';
        // notificar al padre que se ha actualizado algo
        cleanMessage(false);
      }
      // Puedes manejar la respuesta aquí, por ejemplo, mostrar un mensaje de éxito
    } catch (error) {
      console.error('Error al enviar los datos:', error);
      message.text = 'Error al actualizar la contraseña';
      message.status = 'danger';
      cleanMessage(false);
      // Maneja el error (puedes mostrar un mensaje de error en la interfaz de usuario)
    }
  }

  const copyPassword = (event) =>{
    navigator.clipboard.writeText(user.password)
      .then(() => {
        messagePassword = 'Contraseña copiada al portapapeles';
        setTimeout(() => {
          messagePassword = '';
        }, 4300);
      })
      .catch(err => {
        console.error("Error al copiar: ", err);
      });    
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
<form on:submit|preventDefault={handleSave}>
  {#if message.text != ''}
    <div class="alert alert-{message.status}" role="alert">
      {message.text}
    </div>
  {/if}
  
  <div class="row align-items-end">
    <div class="col-4">
      <label class="form-label">Usuario</label>
      <input type="text" bind:value={user.username} class="form-control">
    </div>

    <div class="col-5">
      <label class="form-label">Correo</label>
      <input type="email" bind:value={user.email} class="form-control">
    </div>

    <div class="col-3">
      <button type="submit" class="btn btn-success">
        <i class="fa fa-save"></i> Grabar Usuario
      </button>
    </div>
  </div>
</form>

<div class="col-12 mt-2">
  <div class="mb-3">
    <label class="form-label">Contraseña &nbsp;&nbsp;&nbsp;&nbsp;<span class="text-success">{messagePassword}</span></label>
    <div class="input-group">
      <input type="password" disabled class="form-control" bind:value={user.password}>
      <button type="button" disabled={btnsDisabled} class="btn btn-secondary" on:click={regeneratePassword}><i class="fa fa-random"></i> Regenerar</button>
      <button type="button" disabled={btnsDisabled} class="btn btn-info" on:click={copyPassword}><i class="fa fa-copy"></i> Copiar</button>
      <button type="button" disabled={btnsDisabled} class="btn btn-success" on:click={savePassword}><i class="fa fa-save"></i> Guardar</button>
    </div>
  </div>
</div>

<div class="row g-3 mb-1">
  <!-- Columna 1: Activar Cuenta -->
  <div class="col-md-4">
    <h4 class="subtitle mb-3">Estado de Cuenta</h4>
    <div class="d-flex justify-content-start">
      <button disabled={btnsDisabled} class="btn btn-primary">
        <i class="fa fa-check"></i> Activar Cuenta
      </button>
    </div>
  </div>

  <!-- Columna 2: Enviar Solicitudes -->
  <div class="col-md-8">
    <h4 class="subtitle mb-3">Enviar Solicitudes a Correo</h4>
    <div class="d-flex flex-wrap gap-2">
      <button disabled={btnsDisabled} class="btn btn-info">
        <i class="fa fa-envelope"></i> Activación de Cuenta
      </button>
      <button disabled={btnsDisabled} class="btn btn-warning">
        <i class="fa fa-refresh"></i> Cambio de Contraseña
      </button>
    </div>
  </div>
</div>
