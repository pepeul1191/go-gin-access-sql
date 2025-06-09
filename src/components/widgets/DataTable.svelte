<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import axios from 'axios';
  import { Modal } from 'bootstrap';
  import { navigate } from 'svelte-routing';

  export let fetchURL = null;
  export let deleteURL = fetchURL;
  export let data = [];
  export let columnKeys = [];
  export let columnTypes = [];
  export let columnNames = [];
  export let columnClasses = [];
  export let columnStyles = [];
  export let addButton = {
    display: true,
    disabled: false,
    action: () => {},
  };
  export let actionButtons = [];
  export let pagnation = false;
  export let queryParams = {}

  // delete confirmation modal
  let deleteConfirmationInstance;
  let deleteConfirmationModal;
  let messageConfirmationModal = {
    text: '',
    status: ''
  };
  let idForDeleting = null;
  let btnDisabledDeleteConfirmation = false;

  onMount(() => {
    list();
    deleteConfirmationInstance = new Modal(deleteConfirmationModal);
    console.log(addButton)
  });

  export const askToDeleteRow = (record, key) => {
    idForDeleting = record[key];
    deleteConfirmationInstance.show();
  }

  const cleanMessage = () => {
    setTimeout(() => {
      messageConfirmationModal = {
        text: '',
        status: ''
      };
      btnDisabledDeleteConfirmation = false;
      deleteConfirmationInstance.hide();
    }, 4300);
  }

  export const deleteRowFromDB = () => {
    btnDisabledDeleteConfirmation = true;
    if(deleteURL && idForDeleting){
      axios.delete( // url, data, headers
        deleteURL + '/' + idForDeleting, 
        {
          // params: queryParams,
          headers:{
            //[CSRF.key]: CSRF.value,
          }
        },
      )
      .then((response) => {
        console.log(response);
        data = data.filter(item => item.id !== idForDeleting);
        messageConfirmationModal.text = response.data.message ? response.data.message : 'Registro borrado correctamente';
        messageConfirmationModal.status = 'success';
        cleanMessage();
      })
      .catch((error) => {
        messageConfirmationModal.text = error.message ? error.message : 'Registro borrado correctamente';;
        messageConfirmationModal.status = 'danger';
        console.error(error);
        cleanMessage();
      })
      .then(() => {
        
      });
    }else{
      console.error('No hay URL para eliminar datos o id a eliminar');
    }
  }

  export const addRow = () => {
    alert('addRow');
  }

  export const goToLink = (href) => {
    navigate(href)
  }

  export const goToHref = (href) => {
    window.location.href = href;
  }

  export const openTab = (href) => {
    window.open(href, '_blank');
  }

  export const list = () => {
    if(fetchURL){
      axios.get( // url, data, headers
        fetchURL, 
        {
          params: Object.keys(queryParams).length > 0 ? queryParams : undefined,
          headers:{
            //[CSRF.key]: CSRF.value,
          }
        },
      )
      .then(function (response) {
        data = [];
        data = response.data;
      })
      .catch(function (error) {
        console.error(error);
      })
      .then(function () {
        
      });
    }else{
      console.error('No hay URL para traer datos');
    }
  };
</script>

<style>
  .table-controls > .btn{
    margin-left: 10px;
  }

  .page-link{
    border-radius: 0px !important;
  }

  .tfoot > tr > td {
    padding-left: 0px !important;
  }

  span.page-link{
    color: #343434 !important;
  }

  .page-item{
    margin-left: 0px;
  }

  .text-end > .btn{
    margin-left: 5px;
  }
</style>
<!-- modal -->
<div bind:this={deleteConfirmationModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog modal-dialog-centered">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Confirmación de Eliminación</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        {#if messageConfirmationModal.text != ''}
          <div class="alert alert-{messageConfirmationModal.status}" role="alert">
            {messageConfirmationModal.text}
          </div>
        {/if}
        ¿Seguro que quiere borrar el registro?
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" disabled={btnDisabledDeleteConfirmation}>
          <i class="fa fa-times"></i>Cancelar</button>
        <button type="button" class="btn btn-danger" disabled={btnDisabledDeleteConfirmation}
        on:click={deleteRowFromDB}>
          <i class="fa fa-trash"></i> Eliminar
        </button>
      </div>
    </div>
  </div>
</div>
<!-- Tabla de resultados -->
<div class="d-flex justify-content-between align-items-center">
  <!-- Parte izquierda: Filtro de filas por página -->
  <div class="d-flex align-items-center me-3">
    {#if pagnation}
      <label for="rows-per-page" class="form-label mb-0 me-2">Filas por página:</label>
      <select class="form-select" id="rows-per-page" style="width: 120px;">
        <option value="5">5</option>
        <option value="10">10</option>
        <option value="15">15</option>
        <option value="20">20</option>
      </select>
    {/if}
  </div>
  <!-- Parte derecha: Botón "Agregar Registro" con ícono de Font Awesome -->
  <div class="d-flex gap-2">
    <button
      class="btn btn-primary d-flex align-items-center"
      disabled={addButton.disabled}
      on:click={() => {
        if (typeof addButton.action === 'function') {
          addButton.action();
        } else {
          alert('No se seteado un evento');
        }
      }}
    >
      <i class="fa fa-plus me-2"></i> Agregar Registro
    </button>
    <button class="btn btn-success d-flex align-items-center">
      <i class="fa fa-check me-2"></i> Guardar Cambios
    </button>
  </div>
</div>    
<table class="table table-striped">
  <thead>
    <tr>
      {#each columnNames as key, i}
        <th class="{columnClasses[i]}" style="{columnStyles[i]}" scope="col">{columnNames[i]}</th>
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each data as record}
    <tr>
      {#each columnKeys as key, i}
        <td class="{columnClasses[i]}" style="{columnStyles[i]}">{record[key]}</td>
      {/each}
      {#if actionButtons.length > 0}
        <td class="text-end" styles="">
          {#each actionButtons as button}
            <button class="btn {button.class}" on:click={() => {
              if (typeof button.action === 'function') {
                button.action(record);
              } else {
                alert('No se seteado un evento');
              }
            }}><i class="fa {button.icon}"></i> {button.label}</button>
          {/each}
        </td>
      {/if}
    </tr>
    {/each}
    </tbody>
  {#if pagnation}
    <tfoot>
      <tr>
        <td colspan="6">
          <div class="d-flex justify-content-between align-items-center">
            <!-- Texto con el rango de filas mostradas (izquierda) -->
            <div class="text-left">
              <span>Página 1 de 10 - Mostrando filas 1-10 de 100</span>
            </div>
            <!-- Paginación (derecha) -->
            <nav aria-label="Page navigation">
              <ul class="pagination mb-0">
                <!-- Página Primero -->
                <li class="page-item disabled">
                  <a class="page-link" href="#" tabindex="-1">
                    <i class="fa fa-angle-double-left"></i> Primero
                  </a>
                </li>
                <!-- Página Anterior -->
                <li class="page-item disabled">
                  <a class="page-link" href="#" tabindex="-1">
                    <i class="fa fa-angle-left"></i> Anterior
                  </a>
                </li>
                <!-- Página Siguiente -->
                <li class="page-item">
                  <a class="page-link" href="#">
                    Siguiente <i class="fa fa-angle-right"></i>
                  </a>
                </li>
                <!-- Página Último -->
                <li class="page-item">
                  <a class="page-link" href="#">
                    Último <i class="fa fa-angle-double-right"></i>
                  </a>
                </li>
              </ul>
            </nav>
          </div>
        </td>
      </tr>
    </tfoot>
  {/if}
</table>