import { LoginComponent } from "./login.component"

describe('LoginComponent', () => {
    it('can access the URL', () => {
        // clicking the anchor causes the browser to follow the link
        cy.window()
    })
    it('can submit an empty response', () => {
        // clicking the anchor causes the browser to follow the link
        cy.get('input').submit() // Submit a form
    })
})