-- Write your query below
select employee_id, 
        case 
            when employee_id % 2 <> 0 and substring(name, 1,1) <> 'M' then salary
            else 0
        end as bonus
from employees
order by employee_id