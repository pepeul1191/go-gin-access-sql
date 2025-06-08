<script>
  import DataTable from '../../widgets/DataTable.svelte';
  import { onMount } from 'svelte';
  import { Modal } from 'bootstrap';
  
  let modalInstance;
  let systemDetailModal;
  let message = null;
  let systemDataTable;

  const addSystem = () => {
    modalInstance.show();
  }

  const handleClose = () => {
    console.log('Modal cerrado');
    // Aquí puedes emitir un evento o cambiar estado
  }

  onMount(() => {
    // montar acciones de la tabla
      // ejemplos
      //systemDataTable.dataParams.addButton.action = () => systemDataTable.addRow();
      //systemDataTable.dataParams.addButton.action = () => systemDataTable.goToLink('/users');
      //systemDataTable.dataParams.addButton.action = () => systemDataTable.goToHref(BASE_URL + 'hola');
      //systemDataTable.dataParams.addButton.action = () => systemDataTable.openTab(BASE_URL + 'hola');
    systemDataTable.dataParams.addButton.action = () => addSystem();
    modalInstance = new Modal(systemDetailModal);
    systemDetailModal.addEventListener('hidden.bs.modal', handleClose);
  });
</script>

<style>

</style>

<div bind:this={systemDetailModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Título del Modal</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        <p>Contenido del modal...</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
      </div>
    </div>
  </div>
</div>

<div class="container my-2">
  <div class="row">
    <h1 class="mb-2 subtitle">Gestión de Sistemas</h1>
  </div>
  <hr>
  <div class="row subtitle-row">
    <h4 class="subtitle">Filtros</h4>
  </div>
  {#if message}
    <div class="alert alert-{message.status}" role="alert">
      {message.text}
    </div>
  {/if}
  <div class="container">
    <!-- Formulario de Búsqueda -->
    <form class="mb-4">
        <div class="row">
            <div class="col-md-3">
                <label for="name" class="form-label">Buscar por Nombre</label>
                <input type="text" class="form-control" id="name" placeholder="Nombre">
            </div>
            <div class="col-md-5">
                <label for="description" class="form-label">Buscar por Descripción</label>
                <input type="text" class="form-control" id="description" placeholder="Descripción">
            </div>
            <div class="col-md-4 d-flex align-items-end">
              <button type="submit" class="btn btn-primary me-2">
                <i class="fa fa-search me-2"></i> Buscar
              </button>
              <button type="reset" class="btn btn-secondary">
                <i class="fa fa-eraser me-2"></i> Limpiar
              </button>
            </div>
        </div>
    </form>
  </div>
  <div class="row subtitle-row">
    <h4 class="subtitle">Listado de Sistemas</h4>
  </div>
  <div class="container">
    <DataTable 
      bind:this={systemDataTable} 
      dataParams = {{
        fetchURL: BASE_URL + 'apis/v1/systems',
        columnKeys: ['id', 'name', 'description',],
        columnTypes: ['id', 'td', 'td'],
        columnNames: ['ID', 'Nombre', 'Descripción', 'Acciones'],
        columnClasses: ['d-none', '', '', 'text-end'],
        addButton: {
          display: true,
          disabled: false,
        },
      }}
    />
  </div>
</div>