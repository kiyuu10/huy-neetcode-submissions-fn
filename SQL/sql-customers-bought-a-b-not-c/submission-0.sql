-- Write your query below
select c.customer_id, c.customer_name
from customers as c
join orders as o on c.customer_id = o.customer_id
where o.product_name in('A','B') and c.customer_id not in (select customer_id from orders where product_name='C')
group by c.customer_id, c.customer_name
having count(distinct product_name)=2
order by c.customer_name

