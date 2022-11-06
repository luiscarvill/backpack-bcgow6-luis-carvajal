use empl_dep;
-- 1. Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select e.nombre, e.cargo, d.direccion
	from  empleados e
		join departamentos d on d.id = e.departamentos_id;
-- 2. Visualizar los departamentos con más de cinco empleados.
select d.*, count(d.id) empleados from departamentos d
	left join empleados e on d.id = e.departamentos_id
    group by d.id 
    having empleados > 5
;
-- 3. Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select e.nombre, e.sueldo, d.nombre from empleados e
	join departamentos d on d.id = e.departamentos_id
    where e.cargo in (
		select e2.cargo from empleados e2 where e2.nombre = 'Mito' and e2.apellido = 'Barchuk'
    );
 
 -- 4. Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
 
 select * from empleados e 
 join departamentos d on d.id = e.departamentos_id
 where d.nombre = 'contabilidad'
 order by e.nombre;
 -- 5. Mostrar el nombre del empleado que tiene el salario más bajo.
 
 select e.nombre from empleados e order by e.sueldo asc limit 1;
-- 6. Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
 select e.* from empleados e 
 join departamentos d on d.id = e.departamentos_id
 where d.nombre = 'Ventas'
 order by e.sueldo desc limit 1;
 