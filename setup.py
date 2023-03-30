from setuptools import setup

setup(
    name="kolos",
    version="1.0",
    include_package_data=True,
    install_requires=["pyperclip==1.8.2",
                      "click==8.1.3"],
    py_modules=["main"],
    entry_points="""
        [console_scripts]
        kolos=main:cli
    """,
)
