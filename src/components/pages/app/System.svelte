<script>
  import { onMount } from 'svelte';
  import { Modal } from 'bootstrap';
  import DataTable from '../../widgets/DataTable.svelte';
  import SystemDetail from '../../forms/SystemDetail.svelte';
	import SystemFilters from '../../forms/SystemFilters.svelte';
  
  let modalInstance;
  let systemFormInstance;
  let systemDetailModal;
  let message = null;
  let systemDataTable;
  let modalTitle;

  const addSystem = () => {
    modalTitle = 'Agregar Sistema'
    systemFormInstance.clean();
    modalInstance.show();
  }

  const handleClose = () => {
    // systemDataTable.list();
    // Aquí puedes emitir un evento o cambiar estado
  }

  const handleFormSave = (event) => {
    systemDataTable.list();
    modalInstance.hide();
  };

  const editSystem = (system) => {
    modalTitle = 'Editar Sistema'
    systemFormInstance.clean();
    systemFormInstance.loadSystem(system);
    modalInstance.show();
  }

  const handleSearchFilter = (event) => {
    const { name, description } = event.detail;
    systemDataTable.queryParams = {name,description};
    systemDataTable.list();
  }
  
  const handleCleanFilter= () => {
    systemDataTable.queryParams = {};
    systemDataTable.list();
  }

  onMount(() => {
    // montar acciones de la tabla
      // ejemplos
      //systemDataTable.addButton.action = () => systemDataTable.addRow();
      //systemDataTable.addButton.action = () => systemDataTable.goToLink('/users');
      //systemDataTable.addButton.action = () => systemDataTable.goToHref(BASE_URL + 'hola');
      //systemDataTable.addButton.action = () => systemDataTable.openTab(BASE_URL + 'hola');
    
    modalInstance = new Modal(systemDetailModal);
    systemDetailModal.addEventListener('hidden.bs.modal', handleClose);
    // table action buttons
    systemDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-list',
        label: 'Roles y Permisos',
        action: () => {
          alert('ver');
        }
      },
      {
        class: 'btn-secondary',
        icon: 'fa-users',
        label: 'Usuarios',
        action: () => {
          alert('ver');
        }
      },
      {
        class: 'btn-secondary',
        icon: 'fa-pencil',
        label: 'Editar',
        action: editSystem
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          systemDataTable.askToDeleteRow(record, 'id');
        }
      },
    ];
  });
</script>

<style>

</style>

<div bind:this={systemDetailModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog modal-lg modal-dialog-centered">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">{modalTitle}</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        <SystemDetail 
          bind:this={systemFormInstance} 
          on:saved={handleFormSave} />
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
    <h4 class="subtitle">Filtros de Búsqueda</h4>
  </div>
  {#if message}
    <div class="alert alert-{message.status}" role="alert">
      {message.text}
    </div>
  {/if}
  <div class="container">
    <SystemFilters 
    on:search={handleSearchFilter} 
    on:clean={handleCleanFilter} />
  </div>
  <div class="row subtitle-row">
    <h4 class="subtitle">Listado de Sistemas</h4>
  </div>
  <div class="container">
    <DataTable 
      bind:this={systemDataTable}
      fetchURL={BASE_URL + 'apis/v1/systems'}
      columnKeys={['id', 'name', 'description']}
      columnTypes={['id', 'td', 'td']}
      columnNames={['ID', 'Nombre', 'Descripción', 'Acciones']}
      columnStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 400px;', 'max-width: 150px;']}
      columnClasses={['d-none', '', '', 'text-end']}
      addButton={{
        display: true,
        disabled: false,
        action: addSystem
      }}
      pagination = {{
        display: true,
        step: 10,
        totalPages: 0,
        actualPage: 1
      }}
      actionButtons={[]} 
    />
  </div>
</div>