# Getting Started Mini Project

## API References

<details>
  <summary>Farm API</summary>
  
  **Resource URL**: api/v1/farms

  **Resource Information**:
   1. Reponse Format: JSON
   2. Requires Authentication: No
   3. Rate Limited: No

  <details>
    <summary>GET Methods</summary>
    1. Without parameters.
      - Example URL Request: api/v1/farms
      - Example Response Success:

            {
                "Status": 200,
                "Message": "GetFarm_All",
                "Data": [
                    {
                        "ID": 1,
                        "Name": "Farm 1",
                        "Description": {
                            "String": "a",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "a",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-11 15:41:41",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "2022-06-08 15:44:45",
                            "Valid": true
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                        },
                        {
                        "ID": 2,
                        "Name": "Farm 2",
                        "Description": {
                            "String": "a",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "a",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-11 15:41:41",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "",
                            "Valid": false
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                    },
                    {
                        "ID": 3,
                        "Name": "test",
                        "Description": {
                            "String": "",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "test update",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-12 11:34:52",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "2022-06-12 04:48:45",
                            "Valid": true
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                    },
                    {
                        "ID": 5,
                        "Name": "test create from update",
                        "Description": {
                            "String": "",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "test create from update",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-12 12:00:18",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "",
                            "Valid": false
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                    }
                ]
            }

    2. With Parameters
      - Example URL Request: /api/v1/farms?id=2
      - Example Response Success:

            {
                "Status": 200,
                "Message": "GetFarm_ID",
                "Data": [
                    {
                        "ID": 2,
                        "Name": "Farm 2",
                        "Description": {
                            "String": "a",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "a",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-11 15:41:41",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "",
                            "Valid": false
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                    }
                ]
            }

  </details>
  
  <details>
   <summary>POST Methods</summary>
    - Example URL Request: /api/v1/farms
    - Example Body:

            {
                "Name":"Farm Name 2",
                "Description":"Description Farm Name 2",
                "Thumbnails":"Thumbnails Farm Name 2"
            }

    - Example Response Duplicate:

            {
                "Status": 500,
                "Message": "Error 1062: Duplicate entry 'test' for key 'name'",
                "Data": []
            }

    - Example Response Success:

            {
                "Status": 200,
                "Message": "CreateFarm",
                "Data": null
            }

  </details>

  <details>
   <summary>PUT Methods</summary>
    - Example URL Request: /api/v1/farms
    - Example Body:

            {
                "ID": 2,
                "Name":"Updating Farm Name 2",
                "Description":"Updating Description Farm Name 2",
                "Thumbnails":"Updating Thumbnails Farm Name 2"
            }

    - Example Response Duplicate:

            {
                "Status": 500,
                "Message": "Error 1062: Duplicate entry 'test' for key 'name'",
                "Data": []
            }

    - Example Response Creating:

            {
                "Status": 200,
                "Message": "CreateFarm",
                "Data": null
            }

    - Example Response Updating:

            {
                "Status": 200,
                "Message": "UpdateFarm",
                "Data": null
            }

  </details>

  <details>
   <summary>DELETE Methods</summary>
    - Example URL Request: /api/v1/farms
    - Example Body:

            {
                "ID": 2
            }

    - Example Response Success:

            {
                "Status": 200,
                "Message": "DeleteFarm",
                "Data": null
            }
            
  </details>
</details>

<details>
  <summary>Ponds API</summary>
  
  **Resource URL**: api/v1/ponds

  **Resource Information**:
   1. Reponse Format: JSON
   2. Requires Authentication: No
   3. Rate Limited: No

  <details>
    <summary>GET Methods</summary>
    1. Without parameters.
      - Example URL Request: api/v1/ponds
      - Example Response Success:

                {
                    "Status": 200,
                    "Message": "GetPonds_All",
                    "Data": [
                        {
                            "ID": 1,
                            "Farm_ID": 1,
                            "Name": "Pond 1-1A",
                            "Description": {
                                "String": "",
                                "Valid": true
                            },
                            "Thumbnails": {
                                "String": "",
                                "Valid": true
                            },
                            "Created_at": {
                                "String": "2022-06-11 23:42:03",
                                "Valid": true
                            },
                            "Updated_at": {
                                "String": "",
                                "Valid": false
                            },
                            "Deleted_at": {
                                "String": "",
                                "Valid": false
                            }
                        },
                        {
                            "ID": 2,
                            "Farm_ID": 1,
                            "Name": "Pond 1-1B",
                            "Description": {
                                "String": "",
                                "Valid": true
                            },
                            "Thumbnails": {
                                "String": "",
                                "Valid": true
                            },
                            "Created_at": {
                                "String": "2022-06-11 23:42:03",
                                "Valid": true
                            },
                            "Updated_at": {
                                "String": "",
                                "Valid": false
                            },
                            "Deleted_at": {
                                "String": "",
                                "Valid": false
                            }
                        },
                        {
                            "ID": 7,
                            "Farm_ID": 1,
                            "Name": "test create from update",
                            "Description": {
                                "String": "",
                                "Valid": true
                            },
                            "Thumbnails": {
                                "String": "test create from update",
                                "Valid": true
                            },
                            "Created_at": {
                                "String": "2022-06-12 12:13:21",
                                "Valid": true
                            },
                            "Updated_at": {
                                "String": "",
                                "Valid": false
                            },
                            "Deleted_at": {
                                "String": "",
                                "Valid": false
                            }
                        }
                    ]
                }

    2. With Parameters
      - Example URL Request: /api/v1/ponds?id=2
      - Example Response Success:

            {
                "Status": 200,
                "Message": "GetPonds_ID",
                "Data": [
                    {
                        "ID": 2,
                        "Farm_ID": 1,
                        "Name": "Pond 1-1B",
                        "Description": {
                            "String": "",
                            "Valid": true
                        },
                        "Thumbnails": {
                            "String": "",
                            "Valid": true
                        },
                        "Created_at": {
                            "String": "2022-06-11 23:42:03",
                            "Valid": true
                        },
                        "Updated_at": {
                            "String": "",
                            "Valid": false
                        },
                        "Deleted_at": {
                            "String": "",
                            "Valid": false
                        }
                    }
                ]
            }

  </details>
  
  <details>
   <summary>POST Methods</summary>
    - Example URL Request: /api/v1/ponds
    - Example Body:

            {
                "Farm_ID": 2,
                "Name":"Ponds Name 2",
                "Description":"Description Pond Name 2",
                "Thumbnails":"Thumbnails Pond Name 2"
            }

    - Example Response Duplicate:
    
            {
                "Status": 500,
                "Message": "Error 1062: Duplicate entry 'test' for key 'name'",
                "Data": []
            }

    - Example Response Success:

            {
                "Status": 200,
                "Message": "CreatePonds",
                "Data": null
            }

  </details>

  <details>
   <summary>PUT Methods</summary>
    - Example URL Request: /api/v1/ponds
    - Example Body:

            {
                "ID": 2,
                "Farm_ID": 2,
                "Name":"Updating Ponds Name 2",
                "Description":"Updating Description Pond Name 2",
                "Thumbnails":"Updating Thumbnails Pond Name 2"
            }

    - Example Response Duplicate:

            {
                "Status": 500,
                "Message": "Error 1062: Duplicate entry 'test' for key 'name'",
                "Data": []
            }

    - Example Response Creating:

            {
                "Status": 200,
                "Message": "CreatePonds",
                "Data": null
            }

    - Example Response Updating:

            {
                "Status": 200,
                "Message": "UpdatePonds",
                "Data": null
            }

  </details>

  <details>
   <summary>DELETE Methods</summary>
    - Example URL Request: /api/v1/ponds
    - Example Body:

            {
                "ID": 2
            }

    - Example Response Success:

            {
                "Status": 200,
                "Message": "DeletePonds",
                "Data": null
            }

  </details>

  <details>
   <summary>API Analyst</summary>
    **Resource URL**: api/v1/farms

    **Resource Information**:
     1. Reponse Format: JSON
     2. Requires Authentication: No
     3. Rate Limited: No

     <details>
      <summary>GET Methods</summary>
       - Example URL Request: /api/v1/api_analyst
       - Example Response Success:

            {
                "Status": 200,
                "Message": "GetAPIAnalyst_All",
                "Data": [
                    {
                        "Method": "DELETE",
                        "Path": "/api/v1/farms",
                        "Count": "2",
                        "UA": "Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "DELETE",
                        "Path": "/api/v1/ponds",
                        "Count": "1",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "GET",
                        "Path": "/api/v1/api_analyst",
                        "Count": "3",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "GET",
                        "Path": "/api/v1/farms",
                        "Count": "1",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "GET",
                        "Path": "/api/v1/ponds",
                        "Count": "2",
                        "UA": "Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "GET",
                        "Path": "/api/v1/ponds",
                        "Count": "1",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "POST",
                        "Path": "/api/v1/farms",
                        "Count": "3",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "POST",
                        "Path": "/api/v1/ponds",
                        "Count": "3",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "PUT",
                        "Path": "/api/v1/farms",
                        "Count": "3",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    },
                    {
                        "Method": "PUT",
                        "Path": "/api/v1/ponds",
                        "Count": "2",
                        "UA": "Thunder Client (https://www.thunderclient.com)"
                    }
                ]
            }

     </details>

  </details>

 </details>

</details>


## Database & Storage Systems 
MySQL

### Setup:
1. Download and Install XAMPP / MySQL
2. Make sure there is no database with name "delosaqua"
2. Import sql File in folder resources/databases into MySQL
3. Make sure all Environtment in .env file for your database system

### Running:
1. if you installing XAMPP please Open XAMPP
2. Click Start on Apache Module and MySQL or Run in terminal with command mysql -u userName -p
3. Enter password if you run on terminal


## Application System
Go-lang

### Preparing:
1. Download and Install Go
2. Open Terminal
3. Go to this MiniProject folder
4. type command go run main.go


