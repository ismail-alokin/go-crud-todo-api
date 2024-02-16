module taskSchema {

    type Tasks {
        required title: str;
        description: str;
        due_date: cal::local_date;
        created_at: datetime {
            default := (select datetime_current());
        };
    }

}