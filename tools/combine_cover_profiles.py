
import os

profile_ext = "coverprofile"
master_profile = "MASTER." + profile_ext

def main():
  file_dir = os.path.dirname(os.path.realpath(__file__))
  app_dir  = file_dir.rsplit('/',1)[0]
  print app_dir

  profiles = get_profiles(app_dir)

  f_dir = os.path.join(app_dir, master_profile)
  
  with open(f_dir, "w+") as f:
    f.write("mode: atomic\n")

    for p in profiles:
      with open(p) as temp_f:
        for i in xrange(1):
          temp_f.next()
        for line in temp_f:
          f.write(line)
      os.remove(p)


def get_profiles(app_dir):
  profiles = []

  for root, dirs, files in os.walk(app_dir):
    for f in files:
      if '.' in f:
        if  f.rsplit('.', 1)[1] == profile_ext:
          if f != master_profile:
            profiles.append(os.path.join(root, f))

  return profiles


if __name__ == "__main__":
  main()