
-- 1. Mostrar el título y el nombre del género de todas las series.
select s.title, g.name from series s
	JOIN genres g on s.genre_id = g.id ;
-- 2.Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select e.title, a.first_name, a.last_name from episodes e
	join actor_episode ae on ae.episode_id = e.id 
    join actors a on ae.actor_id = a.id;    
-- 3.Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select s.title, 
	(select s2.number from seasons s2
		where s2.serie_id = s.id order by s2.number desc limit 1) as number
	from series s;
    -- select * from series s join seasons s2 on s2.serie_id = s.id;
-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
-- select *
--  from genres g
-- 	join movies m on m.genre_id = g.id order by g.name;
select g.name, count(g.name) quantity
 from genres g
	join movies m on m.genre_id = g.id 
    group by g.name
    having(quantity) >= 3;
    
-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select a.first_name, a.last_name from actors a
	join actor_movie am on am.actor_id = a.id
    join movies m on am.movie_id = m.id
    where m.title like '%la guerra de las galaxias%'
    group by a.first_name, a.last_name;


		