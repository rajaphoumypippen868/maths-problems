def main():
    import sys
    if len(sys.argv) > 1:
        print("Usage: python3 " + sys.argv[0] + " <problem_name>")
        sys.exit(1)
    else:
        print("Please provide a problem name.")

if __name__ == "__main__":
    main()
