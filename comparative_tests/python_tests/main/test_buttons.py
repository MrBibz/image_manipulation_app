import tkinter as tk
from tkinter import messagebox

def on_button_click(button_name):
    """
    Handle button clicks and display which button was clicked.
    """
    print(f"{button_name} clicked")
    messagebox.showinfo("Button Clicked", f"You clicked {button_name}")

def buttons_test():
    """
    Create a Tkinter window with four buttons.
    """
    # Create the main window
    window = tk.Tk()
    window.title("Test Buttons")
    window.geometry("800x600")

    # Add some padding to the window
    padding = {"padx": 10, "pady": 10}

    # Create buttons and bind their actions
    button1 = tk.Button(window, text="Button 1", command=lambda: on_button_click("Button 1"))
    button1.pack(**padding)

    button2 = tk.Button(window, text="Button 2", command=lambda: on_button_click("Button 2"))
    button2.pack(**padding)

    button3 = tk.Button(window, text="Button 3", command=lambda: on_button_click("Button 3"))
    button3.pack(**padding)

    button4 = tk.Button(window, text="Button 4", command=lambda: on_button_click("Button 4"))
    button4.pack(**padding)

    # Run the Tkinter event loop
    window.mainloop()

if __name__ == "__main__":
    print("\nTest Buttons:")
    buttons_test()
