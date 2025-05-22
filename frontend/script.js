document.querySelectorAll('.grid-button').forEach(button => {
    button.addEventListener('click', function() {
        const tablaContainer = document.getElementById('tablaResultados');
        const buttonId = this.id.replace('btn', '').toLowerCase();
        const loadingId = `loading-${buttonId}`;
        
        // Mostrar loading
        const loadingDiv = document.createElement('div');
        loadingDiv.className = 'loading';
        loadingDiv.textContent = 'Cargando datos...';
        loadingDiv.id = loadingId;
        tablaContainer.appendChild(loadingDiv);
        
        const endpoint = `/${buttonId}`;
        const buttonText = this.textContent.trim();
        
        fetch(endpoint)
            .then(response => {
                if (!response.ok) throw new Error(`Error ${response.status}`);
                return response.json();
            })
            .then(data => {
                if (!data) throw new Error("No se recibieron datos");
                
                // Eliminar el loading
                const loadingToRemove = document.getElementById(loadingId);
                if (loadingToRemove) loadingToRemove.remove();
                
                // Determinar qué tipo de tabla mostrar
                if (buttonText.includes('DÓLAR') && !buttonText.includes('HISTÓRICO')) {
                    crearTablaDolar(data, buttonText);
                } else if (buttonText.includes('HISTÓRICO')) {
                    crearTablaDolarHistorico(data, buttonText);
                } else {
                    crearTablaInstrumentos(data, buttonText);
                }
                
                // Limitar a 9 tablas (3 filas de 3)
                const tablas = document.querySelectorAll('.tabla-wrapper');
                if (tablas.length > 9) {
                    tablaContainer.removeChild(tablas[tablas.length - 1]);
                }
            })
            .catch(error => {
                const loadingToReplace = document.getElementById(loadingId);
                if (loadingToReplace) {
                    loadingToReplace.className = 'error';
                    loadingToReplace.textContent = `Error: ${error.message}`;
                }
                console.error('Error:', error);
            });
    });
});

function crearTablaInstrumentos(data, titulo) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    
    // Título
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    // Contenedor para scroll
    const tableContainer = document.createElement('div');
    tableContainer.className = 'data-table-container';
    
    // Tabla
    const table = document.createElement('table');
    table.className = 'data-table';
    
    // Encabezados
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Instrumento', 'Precio'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    // Cuerpo
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.slice(0, 10).forEach(item => {
        const row = document.createElement('tr');
        
        // Instrumento
        const tdInstrumento = document.createElement('td');
        tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
        row.appendChild(tdInstrumento);
        
        // Precio
        const tdPrecio = document.createElement('td');
        tdPrecio.className = 'price-cell';
        const precio = item.c || item.price || item.precio || item.value;
        tdPrecio.textContent = precio !== undefined ? `$${parseFloat(precio).toFixed(2)}` : 'N/A';
        row.appendChild(tdPrecio);
        
        tbody.appendChild(row);
    });
    
    // Ver más
    if (datos.length > 10) {
        const row = document.createElement('tr');
        row.className = 'ver-mas-row';
        const td = document.createElement('td');
        td.colSpan = 2;
        td.style.textAlign = 'center';
        
        const verMasBtn = document.createElement('span');
        verMasBtn.className = 'ver-mas-btn';
        verMasBtn.textContent = `Ver ${datos.length - 10} más...`;
        verMasBtn.onclick = () => {
            row.remove();
            datos.slice(10).forEach(item => {
                const newRow = document.createElement('tr');
                
                const tdInstrumento = document.createElement('td');
                tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
                newRow.appendChild(tdInstrumento);
                
                const tdPrecio = document.createElement('td');
                tdPrecio.className = 'price-cell';
                const precio = item.c || item.price || item.precio || item.value;
                tdPrecio.textContent = precio !== undefined ? `$${parseFloat(precio).toFixed(2)}` : 'N/A';
                newRow.appendChild(tdPrecio);
                
                tbody.appendChild(newRow);
            });
        };
        
        td.appendChild(verMasBtn);
        row.appendChild(td);
        tbody.appendChild(row);
    }
    
    table.appendChild(tbody);
    tableContainer.appendChild(table);
    tablaWrapper.appendChild(tableContainer);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}

function crearTablaDolar(data, titulo) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    
    // Título
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    // Contenedor para scroll
    const tableContainer = document.createElement('div');
    tableContainer.className = 'data-table-container';
    
    // Tabla
    const table = document.createElement('table');
    table.className = 'data-table';
    
    // Encabezados
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Tipo', 'Compra', 'Venta'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    // Cuerpo
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.forEach(item => {
        const row = document.createElement('tr');
        
        // Tipo
        const tdTipo = document.createElement('td');
        tdTipo.textContent = item.nombre || item.casa || 'N/A';
        row.appendChild(tdTipo);
        
        // Compra
        const tdCompra = document.createElement('td');
        tdCompra.className = 'price-cell';
        tdCompra.textContent = item.compra !== undefined ? `$${item.compra.toFixed(2)}` : 'N/A';
        row.appendChild(tdCompra);
        
        // Venta
        const tdVenta = document.createElement('td');
        tdVenta.className = 'price-cell';
        tdVenta.textContent = item.venta !== undefined ? `$${item.venta.toFixed(2)}` : 'N/A';
        row.appendChild(tdVenta);
        
        tbody.appendChild(row);
    });
    
    table.appendChild(tbody);
    tableContainer.appendChild(table);
    tablaWrapper.appendChild(tableContainer);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}

function crearTablaDolarHistorico(data, titulo) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    
    // Título
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    // Contenedor para scroll
    const tableContainer = document.createElement('div');
    tableContainer.className = 'data-table-container';
    
    // Tabla
    const table = document.createElement('table');
    table.className = 'data-table';
    
    // Encabezados
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Fecha', 'Valor'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    // Cuerpo
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.forEach(item => {
        const row = document.createElement('tr');
        
        // Fecha
        const tdFecha = document.createElement('td');
        if (item.fecha) {
            const fecha = new Date(item.fecha);
            tdFecha.textContent = fecha.toLocaleDateString();
        } else {
            tdFecha.textContent = 'N/A';
        }
        row.appendChild(tdFecha);
        
        // Valor
        const tdValor = document.createElement('td');
        tdValor.className = 'price-cell';
        const valor = item.valor || item.precio || item.c;
        tdValor.textContent = valor !== undefined ? `$${parseFloat(valor).toFixed(2)}` : 'N/A';
        row.appendChild(tdValor);
        
        tbody.appendChild(row);
    });
    
    table.appendChild(tbody);
    tableContainer.appendChild(table);
    tablaWrapper.appendChild(tableContainer);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}