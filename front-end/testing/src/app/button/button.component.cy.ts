import { ButtonComponent } from "./button.component"

describe('ButtonComponent', () => {
    it('can mount', () => {
        cy.mount(ButtonComponent)
    })

    it('can display label in button', () => {
        cy.mount(ButtonComponent, {
            componentProperties: {
                label: 'hello angular'
            }
        })
    })
})