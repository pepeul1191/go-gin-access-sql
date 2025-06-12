<script>
  import { onMount } from 'svelte';
  import DataTable from '../widgets/DataTable.svelte';

  export let systemId = null;

  let alertMessage = {
    text: '',
    status: ''
  };
  let rolePermissionDataTable;

  export const setSystemId = (system) => {
    console.log(system)
    systemId = system.id;
    rolePermissionDataTable.fetchURL = BASE_URL + 'apis/v1/systems/' + systemId + '/roles';
    rolePermissionDataTable.list();
  }

  const addRole = () => {
    rolePermissionDataTable.addRow();
  }

  const handleTableAlert = (callback) => { 
    alertMessage = callback.detail;
    console.log(callback.detail)
    setTimeout(() => {
      alertMessage = {
        text: '',
        status: '',
      };
    }, 4300);
  }

  onMount(() => {
    // montar acciones de la tabla
      // ejemplos
    rolePermissionDataTable.addButton.action = () => rolePermissionDataTable.addRow();
    //systemDataTable.addButton.action = () => systemDataTable.goToLink('/users');
    //systemDataTable.addButton.action = () => systemDataTable.goToHref(BASE_URL + 'hola');
    //systemDataTable.addButton.action = () => systemDataTable.openTab(BASE_URL + 'hola');
    // table action buttons
    rolePermissionDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-list',
        label: 'Permisos',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          //console.log(record);
        }
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          //console.log(record);
          rolePermissionDataTable.deleteRow(record, 'id');
        }
      },
    ];
  });
</script>

<style></style>

{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}

<div class="row g-2">
  <div class="col-md-6">
    <div class="row subtitle-row">
      <h4 class="subtitle">Lista de Roles</h4>
    </div>
    <DataTable 
      bind:this={rolePermissionDataTable}
      fetchURL={BASE_URL + 'apis/v1/systems/' + systemId + '/roles'}
      saveURL={BASE_URL + 'apis/v1/roles/' + systemId}
      columnKeys={['id', 'name', ]}
      columnTypes={['id', 'input[text]', ]}
      columnNames={['ID', 'Nombre', 'Acciones']}
      columnStyles={['max-width: 50px;', 'max-width: 150px;', 'max-width: 150px;']}
      columnClasses={['d-none', '', 'text-end']}
      messages = {{
        success: 'Datos actualizados', 
        errorNetwork: 'No se pudo listar los roles del sistema. No hay conexión con el servidor.',
        notFound: 'No se pudo listar los roles del sistema. Recurso no encontrado.',
        serverError:'No se pudo listar los roles del sistema. Error interno del servidor',
        requestError: 'No se pudo listar los roles del sistema. No se recibió respuesta del servidor',
        otherError: 'No se pudo listar los roles del sistema. Ocurrió un error no esperado al traer los datos del servidor',
      }}
      addButton={{
        display: true,
        disabled: false,
        action: null
      }}
      saveButton={{
        display: true,
        disabled: false,
        action: null
      }}
      pagination = {{
        display: false,
        step: 10,
        totalPages: 0,
        actualPage: 1
      }}
      actionButtons={[]} 
      on:alert={handleTableAlert}
    />
  </div>
  <div class="col-md-6">
    <div class="row subtitle-row">
      <h4 class="subtitle">Lista de Permisos del Rol</h4>
    </div>
  </div>
</div>