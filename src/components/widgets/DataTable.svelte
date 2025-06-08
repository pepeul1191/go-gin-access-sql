<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import axios from 'axios';
  import { navigate } from 'svelte-routing';

  export let data = [];
  export let dataParams = {
    fetchURL: '',
    columnKeys: [],
    columnTypes: [],
    columnNames: [],
    columnClasses: [],
    columnStyles: [],
    addButton: {
      display: true,
      disabled: false,
      action: () => {
        alert()
      }
    },
  };
  export let pagnation = false; 

  onMount(() => {
    list();
  });

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
    axios.get( // url, data, headers
      dataParams.fetchURL, 
      {
        // params: queryParams,
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
</style>

   <!-- Tabla de resultados -->
<div class="d-flex justify-content-between align-items-center">
  <!-- Parte izquierda: Filtro de filas por página -->
  {#if pagnation}
    <div class="d-flex align-items-center me-3">
      <label for="rows-per-page" class="form-label mb-0 me-2">Filas por página:</label>
      <select class="form-select" id="rows-per-page" style="width: 120px;">
        <option value="5">5</option>
        <option value="10">10</option>
        <option value="15">15</option>
        <option value="20">20</option>
      </select>
    </div>
  {:else}
    <div class="d-flex align-items-center me-3"></div>
  {/if}
  <!-- Parte derecha: Botón "Agregar Registro" con ícono de Font Awesome -->
  <div class="d-flex gap-2">
    <button
      class="btn btn-primary d-flex align-items-center"
      disabled={dataParams.addButton.disabled}
      on:click={() => (dataParams.addButton.action ?? alert('No ha seteado un evento'))()}
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
      {#each dataParams.columnNames as key, i}
        <th class="{dataParams.columnClasses[i]}" scope="col">{dataParams.columnNames[i]}</th>
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each data as record}
    <tr>
      {#each dataParams.columnKeys as key, i}
        <td class="{dataParams.columnClasses[i]}">{record[key]}</td>
      {/each}
      <td class="text-end">
        <button class="btn btn-info"><i class="fa fa-eye"></i> Ver</button>
        <button class="btn btn-warning"><i class="fa fa-pencil"></i> Editar</button>
        <button class="btn btn-danger"><i class="fa fa-trash"></i> Eliminar</button>
      </td>
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