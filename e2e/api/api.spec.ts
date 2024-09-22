import { test, expect} from '@playwright/test'
import { deleteData, insertData } from '../database/postgres';

test.beforeAll(insertData);

test.afterAll(deleteData);

test.describe('GET /api/v1/skills/:key', ()=> {
    test('should response a skill with status success when have key', async ({
        request,
    }) => {
        const res = await request.get(`/api/v1/skills/go`)
    
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
        expect.objectContaining({
            "status": "success",
            "data": {
                "key":"go",
                "name": expect.any(String),
                "description": expect.any(String),
                "logo": expect.any(String),
                "tags":expect.any(Array<String>)
            }
        }))
    })

    test('should response a skill with status 500 when key is not exist', async ({
        request,
    }) => {
        const res = await request.get(`/api/v1/skills/java`)
    
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
        expect.objectContaining({
            "status": "error",
            "message": "Skill not found"
        })
    )})
})

test.describe('GET /api/v1/skills', () => {
    test('should response a skill with status success', async({
        request
    }) => {
        const res = await request.get('/api/v1/skills')
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": expect.arrayContaining([
                    expect.objectContaining({
                        "key":expect.any(String),
                        "name": expect.any(String),
                        "description": expect.any(String),
                        "logo": expect.any(String),
                        "tags":expect.any(Array<String>)
                    })
                ])
        })) 
    })
})

test.describe('POST /api/v1/skills', () => {
    test('should response a skill with status created ', async({
        request
    }) => {
        const res = await request.post('/api/v1/skills', 
            {
                data:{
                    key: "python",
                    name: "Python",
                    description: "Python is an interpreted, high-level, general-purpose programming language.",
                    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
                    tags: ["programming language", "scripting"]
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
        expect.objectContaining({
            "status": "success",
            "data": {
                "key":"python",
                "name": "Python",
                "description": expect.any(String),
                "logo": expect.any(String),
                "tags":expect.any(Array<String>)
            }
        })
        )
    })

    test('should response a skill with status bad request when key is already exist', async({
        request
    }) => {
        const res = await request.post('/api/v1/skills', 
            {
                data:{
                    key: "go",
                    name: "Python",
                    description: "Python is an interpreted, high-level, general-purpose programming language.",
                    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
                    tags: ["programming language", "scripting"]
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
        expect.objectContaining({
            "status": "success",
            "data": {
                "key":"go",
                "name": "Python",
                "description": expect.any(String),
                "logo": expect.any(String),
                "tags":expect.any(Array<String>)
            }
        }))
    })
})

test.describe('PUT /api/v1/skills/:key', () => {
    test('should response a skill with status created', async({
        request
    }) => {
        const res = await request.put(`/api/v1/skills/go`,
            {
                data: {
                    name: "go2",
                    description: "go2 is the latest version of Python programming language.",
                    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
                    tags: ["data"]
                }  
            }
        )
        console.log(res.status())
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": {
                    "key":"go",
                    "name": "go2",
                    "description": "go2 is the latest version of Python programming language.",
                    "logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
                    "tags": ["data"]
                }
            })
        )
    })

    test('should response a skill with status conflict', async({
        request
    }) => {
        const res = await request.put(`/api/v1/skills/kotlin`,
            {
                data: {
                    name: "go2",
                    description: "go2 is the latest version of Python programming language.",
                    logo: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
                    tags: ["data"]
                }  
            }
        )
        expect(res.ok()).not.toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        )
    })
})

test.describe('PATCH /api/v1/skills/:key/actions/name',() => {
    test('should response updated skill name with status ok', async({request}) => {
        const res = await request.patch('/api/v1/skills/go/actions/name',
            {
                data: {
                    name:"golang patch"
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": {
                    "key":"go",
                    "name": "golang patch",
                    "description": expect.any(String),
                    "logo": expect.any(String),
                    "tags": expect.any(Array<String>)
                }
            })
        )
    })

    test('should response a skill with status bad request when key is not exists', async({
        request
    }) => {
        const res = await request.patch(`/api/v1/skills/kotlin/actions/name`,
            {
                data: {
                    name: "kotlin",
                }  
            }
        )
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        )
    })
})

test.describe('PATCH /api/v1/skills/:key/actions/description',() => {
    test('should response updated skill description with status ok', async({request}) => {
        const res = await request.patch('/api/v1/skills/go/actions/description',
            {
                data: {
                    description:"golang description"
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": {
                    "key":"go",
                    "name": expect.any(String),
                    "description": "golang description",
                    "logo": expect.any(String),
                    "tags": expect.any(Array<String>)
                }
            })
        )
    })

    test('should response a skill with status bad request when key is not exists', async({
        request
    }) => {
        const res = await request.patch(`/api/v1/skills/kotlin/actions/description`,
            {
                data: {
                    description: "kotlin",
                }  
            }
        )
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        ) 
    })
})
	
test.describe('PATCH /api/v1/skills/:key/actions/logo',() => {
    test('should response updated skill logo with status ok', async({request}) => {
        const res = await request.patch('/api/v1/skills/go/actions/logo',
            {
                data: {
                    logo:"newlogoPath"
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": {
                    "key":"go",
                    "name": expect.any(String),
                    "description": expect.any(String),
                    "logo": "newlogoPath",
                    "tags": expect.any(Array<String>)
                }
            })
        )
    })

    test('should response updated skill logo with status bad request when key is no exists', async({request}) => {
        const res = await request.patch('/api/v1/skills/kotlin/actions/logo',
            {
                data: {
                    logo:"newLogoPath"
                }
            })
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        ) 
    })
})

test.describe('PATCH /api/v1/skills/:key/actions/tags',() => {
    test('should response updated skill tags with status ok', async({request}) => {
        const res = await request.patch('/api/v1/skills/go/actions/tags',
            {
                data: {
                    tags:["test tags"]
                }
            })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": {
                    "key":"go",
                    "name": expect.any(String),
                    "description": expect.any(String),
                    "logo": expect.any(String),
                    "tags": ["test tags"]
                }
            })
        )
    })
    
    test('should response a skill with status bad request when key is not exists', async({request}) => {
        const res = await request.patch('/api/v1/skills/kotlin/actions/tags',
            {
                data: {
                    tags:["test tags"]
                }
            })
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        )
    })
})

test.describe('DELETE /api/v1/skills/:key',() => {
    test('should response message delete completed with status ok', async({request}) => {
        const res = await request.delete('/api/v1/skills/go')
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "Skill deleted"
            })
        )
    })


    test('should response message delete error with status bad request when key is not exists', async({request}) => {
        const res = await request.delete('/api/v1/skills/kotlin')
        expect(res).not.toBeOK()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "error",
                "message": expect.any(String),
            })
        )
    })
    
    
})